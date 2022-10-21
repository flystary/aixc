package net

import (
	"os"
	"github.com/olekukonko/tablewriter"
)

// Ucpe 一
type Ucpe struct {
	sn          string
	model       string
	version     string
	updatetime  string
	masterpopip string
	mastercpeip string
	backuppopip string
	backupcpeip string
	// remoteport  string
}

// Ucpes 多
type Ucpes []Ucpe

var (
	table *tablewriter.Table
)

func init() {
	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
}

// Show UCPE
func (u *Ucpe) Show() {
	table.Append([]string{
		u.sn,
		u.model,
		u.version,
		u.updatetime,
		u.masterpopip,
		u.mastercpeip,
		u.backuppopip,
		u.backupcpeip,
	})
	table.Render()
}

// Show UCPES
func (us *Ucpes) Show() {
	for _, u := range *us {
		table.Append([]string{ u.sn, u.model, u.version,u.updatetime,
			u.masterpopip, u.mastercpeip, u.backuppopip, u.backupcpeip,
		})
	}
	table.Render()
}
