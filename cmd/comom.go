package cmd

import (
	"strconv"
	"strings"
)

type CLI struct {
	sns     Sns
	options *Options
	mode    string
	write   Write
	isSeven bool
	isOk    bool
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
	isEnte     bool
	enterprise string
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

func (s Sns) Check() Sns {
	var sns = make([]string, 0)
	for _, sn := range s {
		sns = append(sns, sn)
	}
	return sns
}

func (c *CLI) CheckMode() *CLI {
	mode := c.mode
	if mode == "valor" || mode == "nexus" || mode == "watsons" || mode == "watsonsha" || mode == "tassadar" || mode == "" {
		c.isOk = true
	}
	c.isOk = false
	return c
}

func (op *Options) Select(opt string, data []string) *Options {

	if opt == "model" {
		op.isModel = true
		op.model = data[0]
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
		op.isEnte = true
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

func (op *Options) isEnterprise(enterprise string) *Options {
	op.enterprise = enterprise
	op.isEnte = true
	return op
}

func (w *Write) Decode(wr string) *Write {
	if wr == "" {
		w.isWrite = false
		return w
	}
	w.isWrite = true
	if strings.Contains(wr, "/") {
		// PATH 路径
		strs := strings.Split(wr, "/")
		lens := len(strs)
		w.name = strs[lens-1]
		w.path = strings.Join(strs[0:lens], "/")
	} else {
		w.name = wr
		w.path = "./"
	}
	// Format 文件格式
	strs := strings.Split(w.name, ".")
	lens := len(strs)
	w.format = strs[lens-1]
	return w
}

type Cmd interface {
	Init
	search()
	show()
	writef()
	conn()
	run()
}

type Init interface {
	Resolv
	initCpe()
	initPop()
	initDve()
}

type Resolv interface {
}

func (c CLI) initCpe() {}
func (c CLI) initPop() {}
func (c CLI) initDve() {}

func (c CLI) run() {
	c.search()
	c.show()
	if !c.write.isWrite {
		c.show()
	}
	c.writef()

}
func (c CLI) search() {}
func (c CLI) show()   {}
func (c CLI) writef() {}
func (c CLI) conn()   {}
