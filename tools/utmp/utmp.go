package utmp

import (
	"encoding/binary"
	"io"
	"os"
)

const (
	UT_LINESIZE = 32
	UT_NAMESIZE = 32
	UT_HOSTSIZE = 256
)

type ExitStatus struct {
	Termination int16
	Exit        int16
}

type TimeVal struct {
	Sec  int32
	Usec int32
}

type Utmp struct {
	// https://man7.org/linux/man-pages/man5/utmp.5.html
	Type     int16             /* Type of record */
	_        [2]byte           /* 内存对齐,不可少*/
	Pid      int32             /* PID of login process */
	Device   [UT_LINESIZE]byte /* Device name of tty - "/dev/" */
	Id       [4]byte           /* Terminal name suffix, or inittab(5) ID */
	User     [UT_NAMESIZE]byte /* Username */
	Host     [UT_HOSTSIZE]byte /* Hostname for remote login, or kernel version for run-level messages */
	Exit     ExitStatus        /* Exit status of a process marked as DEAD_PROCESS; not used by Linux init(1) */
	Session  int32             /* Session ID (getsid(2)), used for windowing */
	Time     TimeVal           /* Time entry was made */
	AddrV6   [16]byte          /* Internet address of remote host; IPv4 address uses just ut_addr_v6[0] */
	Reserved [20]byte
}

func LoadUtmp() []*Utmp {
	file, err := os.Open("/var/run/utmp")
	if err != nil {
		panic("open /var/run/tmp failed: %s" + err.Error())
	}
	defer file.Close()

	var utmps = make([]*Utmp, 0)
	for {
		utmp := new(Utmp)
		err = binary.Read(file, binary.LittleEndian, utmp)
		if err == io.EOF {
			break
		} else if err != nil {
			panic("read /var/run/tmp failed: " + err.Error())
		}
		// #define USER_PROCESS	7	/* Normal process.  */
		if utmp.Type != 7 || utmp.User[0] == 0 {
			// 忽略特殊用户,参考 procps-ng/contrib/utmp.c 源码
			continue
		}
		utmps = append(utmps, utmp)
	}
	return utmps
}
