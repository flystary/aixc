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

// Shower Show
type Shower interface {
	Show()
}

// Conner Conn
// type Conner interface {
// 	MasterIsEmpty() bool
// 	BackupIsEmpty() bool
// 	ConnMaster()
// 	ConnBackup()
// }

var (
	table *tablewriter.Table
	ucpes = make([][]string, 0)
	slices  = make([]string, 0)
)

func init() {
	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
}

// Show UCPE
func (u *Ucpe) Show() {
	ucpes = append(ucpes, []string{
		u.sn,
		u.model,
		u.version,
		u.updatetime,
		u.masterpopip,
		u.mastercpeip,
		u.backuppopip,
		u.backupcpeip,
		},
	)
	table.AppendBulk(ucpes)
	table.Render()
}

// Show UCPES
func (us *Ucpes) Show() {
	for _, u := range *us {
		x := []string{ u.sn, u.model, u.version,u.updatetime,
			u.masterpopip, u.mastercpeip, u.backuppopip, u.backupcpeip,
		}
		ucpes = append(ucpes, x)
	}
	table.AppendBulk(ucpes)
	table.Render()
}
