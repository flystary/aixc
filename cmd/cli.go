package cmd

type CLI struct {
	sns		Sns
	options	Options
	mode	string
	write	Write
	isSeven bool
	isOk	bool
}

type Sns []string

type Options struct {
	model Model
	version Version
	pop Pop
	enterprise Enterprise
	port Port
	update Update
}

type Model struct {
	IsExist bool
	Model string
}

type Version struct {
	IsExist bool
	Version string
}

type Pop struct {
	IsExist bool
	Pop string
}

type Enterprise struct {
	IsExist bool
	Enterprise string
}

type Port struct {
	IsExist bool
	Port int64
}

type Update struct {
	IsExist bool
	Update string
}

type Write struct {
	name	string
	format  string
	path	string
}

func (s Sns) Check() Sns {
	var sns = make([]string, 0)
	for _, sn := range s {
		sns = append(sns, sn)
	}
	return sns
}

func (c *CLI)CheckMode() *CLI {
	mode := c.mode
	if mode == "valor" || mode == "nexus" || mode == "watsons" || mode == "watsonsha" || mode == "tassadar" || mode == "" {
		c.isOk = true
	}
	c.isOk = false
	return c
}

func (o *Options)Decode(op string) *Options {
	return o
}

func (o *Options)isEnterprise(enterprise string) *Options {
	o.enterprise.Enterprise = enterprise
	o.enterprise.IsExist = true
	return o
}

func (w *Write)Decode(wr string) *Write {
	return w
}

func (cli CLI) Search() {}
func (cli CLI) Run() {}
func (cli CLI) Conn() {}
func (cli CLI) Write() {}