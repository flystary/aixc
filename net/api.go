package net


import (
	"os"
	"fmt"
	"github.com/olekukonko/tablewriter"
)


//SearchSeven 6.x
func SearchSeven(sn string) {
	// 已知mode数据6.x版本
	m := snInSevenMode(sn)
	fmt.Printf("sevencpemode: %s\n", m)
	if m == "unknown" {
		os.Exit(13)
	}
	isSnInMode(sn, m)
	model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime := getCpePop(m)
	tableData := [][]string{
		{sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	for _, v := range tableData {
		table.Append(v)
	}
	table.Render()

	getDvc(sn, m)
	fmt.Printf("%v\n", vd)
}

// SearchSevenMany 6.x
func SearchSevenMany(snMany []string) {
	// 已知mode数据6.x版本
	var mode string
	// table
	var tableData = make([][]string, 0)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = snInSevenMode(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}
	fmt.Printf("sevencpemode: %s\n", mode)

	for _, sn := range snMany {
		tables := make([]string, 0)
		isSnInMode(sn, mode)
		model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime := getCpePop(mode)
		tableData = append(tableData, append(tables, sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip))
	}

	table.AppendBulk(tableData)
	table.Render()
}

//Search 6.x/7.x
func Search(sn string) {
	// 多线程查询
	m := getSnInMode(sn)
	fmt.Printf("cpemode: %s\n", m)
	if m == "unknown" {
		os.Exit(14)
	}
	// isSnInMode(sn, m)
	model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime := getCpePop(m)
	tableData := [][]string{
		{sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	for _, v := range tableData {
		table.Append(v)
	}
	table.Render()

	getDvc(sn, m)
	fmt.Printf("%v\n", vd)
	// fmt.Printf("%v\n", vd.ID)
	// fmt.Printf("%v\n", vd.ServerPort)
}

// SearchMany 6.x/7.x
func SearchMany(snMany []string) {
	var mode string
	// table
	var tableData = make([][]string, 0)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = getSnInMode(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}
	fmt.Printf("cpemode: %s\n", mode)

	for _, sn := range snMany {
		tables := make([]string, 0)
		isSnInMode(sn, mode)
		model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime := getCpePop(mode)
		tableData = append(tableData, append(tables, sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip))
	}

	// for _, v := range tableData {
	// 	table.Append(v)
	// }
	table.AppendBulk(tableData)
	table.Render()
}