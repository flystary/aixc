package conn

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	expect "github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// JumpStep
type JumpStep struct {
	Expect  string // 期望看到的字符串（正则）
	Send    string // 匹配后发送的命令
	Timeout time.Duration
}

const (
	timeout = 10 * time.Minute
)

var (
	addr = flag.String("address", "", "address of ssh server")
	user = flag.String("user", "user", "username to use")
	pass = flag.String("pass", "pass", "password to use")
)

func main() {
	flag.Parse()
	if *addr == "" || *pass == "" {
		log.Fatal("必须提供 -address 和 -pass 参数")
	}

	config := &ssh.ClientConfig{
		User:            *user,
		Auth:            []ssh.AuthMethod{ssh.Password(*pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}

	client, err := ssh.Dial("tcp", *addr, config)
	if err != nil {
		log.Fatalf("连接Ubuntu失败: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("创建会话失败: %v", err)
	}
	defer session.Close()

	// 设置本地终端为 Raw 模式 (允许交互)
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(fd, oldState)

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		term.Restore(fd, oldState)
		os.Exit(1)
	}()

	w, h, _ := term.GetSize(fd)
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.ICRNL:         1,
		ssh.ONLCR:         1,
		ssh.TTY_OP_ISPEED: 115200, // 输入速度
		ssh.TTY_OP_OSPEED: 115200, // 输出速度
	}
	if err := session.RequestPty("xterm-256color", w, h, modes); err != nil {
		log.Fatalf("Request PTY Failed: %v", err)
	}

	out, err := session.StdoutPipe() // 远程输出 (Reader)
	if err != nil {
		log.Fatalf("获取 StdoutPipe 失败: %v", err)
	}
	in, err := session.StdinPipe() // 远程输入 (Writer)
	if err != nil {
		log.Fatalf("获取 StdinPipe 失败: %v", err)
	}

	if err := session.Shell(); err != nil {
		term.Restore(fd, oldState)
		log.Fatal(err)
	}

	if in == nil || out == nil {
		log.Fatal("致命错误: 管道未初始化 (sr 或 sw 为 nil)，请检查 SSH 连接或变量作用域")
	}

	// 注意这里的字段名：Logger (双g), In (Reader), Out (Writer)
	exp, _, err := expect.SpawnGeneric(&expect.GenOptions{
		In:    in,
		Out:   out,
		Wait:  func() error { select {} },
		Close: func() error { return nil },
		Check: func() bool { return true },
		// Logger: os.Stdout,
	}, -1, expect.Verbose(true), expect.VerboseWriter(os.Stdout))
	if err != nil {
		term.Restore(fd, oldState)
		log.Fatalf("Spawn 失败: %v", err)
	}
	steps := []JumpStep{
		{Expect: `[#\$]\s$`, Send: "ls\r", Timeout: 2 * time.Second},
		// {Expect: `(?i)password:`, Send: "Second_Hop_Password", Timeout: 5 * time.Second},
		// {Expect: `admin@jump-box:.*[#\$]`, Send: "ls -lh", Timeout: 5 * time.Second},
	}
	if err := runJumpSequence(exp, steps); err != nil {
		term.Restore(fd, oldState)
		// 这里加上回车换行，因为在 Raw 模式下 \n 不会回到行首
		fmt.Printf("\r\n[Error] 自动化跳转失败: %v\r\n", err)
		return
	}
	fmt.Print("\r\n[Info] 自动化完成，进入交互模式...\r\n")

	done := make(chan struct{})
	go func() {
		io.Copy(in, os.Stdin)
		done <- struct{}{}
	}()
	go func() {
		io.Copy(os.Stdout, out)
		done <- struct{}{}
	}()

	<-done
}

func runJumpSequence(exp *expect.GExpect, steps []JumpStep) error {
	for _, step := range steps {
		re := regexp.MustCompile(step.Expect)
		_, _, err := exp.Expect(re, step.Timeout)
		if err != nil {
			return fmt.Errorf("等待模式 '%s' 超时", step.Expect)
		}
		time.Sleep(50 * time.Millisecond)
		// 发送命令并带上回车
		if err := exp.Send(step.Send + "\n"); err != nil {
			return err
		}
	}

	/*
		// C -> D
		exp.Expect(regexp.MustCompile(`C-host`), 10*time.Second)
		exp.Send("ssh root@D_IP\r")

		// D 密码
		exp.Expect(regexp.MustCompile(`password:`), 5*time.Second)
		exp.Send("D_Password\r")

		// D -> E
		exp.Expect(regexp.MustCompile(`D-host`), 10*time.Second)
		exp.Send("ssh root@E_IP\r")

		// E 密码
		exp.Expect(regexp.MustCompile(`password:`), 5*time.Second)
		exp.Send("E_Password\r")

		// 确认最后一步
		exp.Expect(regexp.MustCompile(`E-host`), 5*time.Second)
	*/
	return nil
}
