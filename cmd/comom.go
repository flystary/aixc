package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	sns   Sns
	mode  string
	seven bool
	isOk  bool
	*Options
	*Write
}

type Sns []string

type Options struct {
	*Model
	*Version
	*Pop
	*Enterprise
	*Port
	*Update
}

type Model struct {
	isModel bool
	model   string
}

type Version struct {
	isVersion bool
	version   string
}

type Pop struct {
	isPop bool
	pop   string
}

type Enterprise struct {
	isEnterprise bool
	enterprise   string
}

type Port struct {
	isPort bool
	port   int
}

type Update struct {
	isUpdate bool
	update   string
}

type Write struct {
	isWrite bool
	name    string
	format  string
	path    string
}

func (c *CLI) CheckSns() *CLI {
	var sns = make([]string, 0)
	for _, sn := range c.sns {
		sns = append(sns, sn)
	}
	c.sns = sns
	return c
}

func (c *CLI) CheckMode() *CLI {
	mode := c.mode
	if mode == "valor" || mode == "nexus" || mode == "watsons" || mode == "watsonsha" || mode == "tassadar" || mode == "" {
		c.isOk = true
	}
	c.isOk = false
	return c
}

func (op *Options) SelectOpt(opt string, data []string) *Options {

	if opt == "model" {
		op.isModel = true
		op.model = data[0]
		return op
	}

	if opt == "version" {
		op.isVersion = true
		op.version = data[0]
	}

	if opt == "pop" {
		op.isPop = true
		op.pop = data[0]
	}

	if opt == "enterprise" {
		op.isEnterprise = true
		op.enterprise = data[0]
	}

	if opt == "port" {
		op.isPort = true
		p, _ := strconv.Atoi(data[0])
		op.port = p
	}

	if opt == "update" {
		op.isUpdate = true
		op.update = data[0]
	}

	return op
}

func (w *Write) DecodeWrite(wr string) *Write {
	if wr == "" {
		w.isWrite = false
		return w
	}
	w.isWrite = true
	if strings.Contains(wr, "/") {
		// PATH
		strs := strings.Split(wr, "/")
		lens := len(strs)
		w.name = strs[lens-1]
		w.path = strings.Join(strs[0:lens], "/")
	} else {
		w.name = wr
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("getwd error: %v", err)
			return w
		}
		w.path = pwd
	}
	// txt
	strs := strings.Split(w.name, ".")
	lens := len(strs)
	w.format = strs[lens-1]
	return w
}

type Cmd interface {
	Init
	check()
	search()
	show()
	writef()
	conn()
	run()
}

type Init interface {
	Resolv
	initCPE()
	initPOP()
	initDVE()
}

type Resolv interface {
}

func (c *CLI) initCPE() {
	fmt.Println("init CPE")
}
func (c *CLI) initPOP() {
	fmt.Println("init POP")
}
func (c *CLI) initDVE() {
	fmt.Println("init DVE")
}

func (c *CLI) check() {
	c.CheckMode()
	c.CheckSns()
	c.isOk = true

}

func (c *CLI) run() {
	c.check()

	if c.isOk {
		fmt.Println("check ok!")
	} else {
		return
	}

	fmt.Printf("cli: %+v\n", c)
	fmt.Printf("options.Model: %+#v\n", c.Options.Model)
	fmt.Printf("options.Version: %+#v\n", c.Options.Version)
	fmt.Printf("options.Pop: %+#v\n", c.Options.Pop)
	fmt.Printf("options.Enterprise: %+#v\n", c.Options.Enterprise)
	fmt.Printf("options.Port: %+#v\n", c.Options.Port)
	fmt.Printf("options.Update: %+#v\n", c.Options.Update)
	fmt.Printf("write: %##v\n", c.Write)
	c.search()
	c.show()
	if !c.isWrite {
		c.show()
	}
	c.writef()

}

func (c *CLI) search() {
}
func (c *CLI) show()   {}
func (c *CLI) writef() {}
func (c *CLI) conn()   {}
