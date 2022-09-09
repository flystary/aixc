package net


import (
	"os"
	"fmt"
	"github.com/olekukonko/tablewriter"
)

var (
	mode	string


	model	string
	version	string
	mastercpeip	string
	masterpopip	string
	backupcpeip	string
	backuppopip string
	updatetime	string
)

//SearchSeven 6.x
func SearchSeven(sn string) {
	mode = snInSevenMode(sn)
	fmt.Printf("sevencpemode: %s\n", mode)
	if mode == "unknown" {
		os.Exit(13)
	}
	syncDataMemorybyMode(mode)
	switch mode {
		case "valor":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyValor(sn)
		}
		case "nexus":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyNexus(sn)
		}
		case "watsons":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsons(sn)
		}
		case "watsons_ha":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyZeratul(sn)
		}
	}
	tableData := [][]string{
		{sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	for _, v := range tableData {
		table.Append(v)
	}
	table.Render()
}

// SearchSevenMany 6.x
func SearchSevenMany(snMany []string) {
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
	// 同步数据到内存
	syncDataMemorybyMode(mode)
	for _, sn := range snMany {
		tables := make([]string, 0)
		switch mode {
			case "valor":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyValor(sn)
			}
			case "nexus":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyNexus(sn)
			}
			case "watsons":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsons(sn)
			}
			case "watsons_ha":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsonsHa(sn)
			}
			case "tassadar":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyZeratul(sn)
			}
		}
		tableData = append(tableData, append(tables, sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip))
	}
	table.AppendBulk(tableData)
	table.Render()
}

//Search 6.x/7.x
func Search(sn string) {
	// 多线程查询
	mode = threadQueryMode(sn)
	fmt.Printf("cpemode: %s\n", mode)
	if mode == "unknown" {
		os.Exit(14)
	}
	switch mode {
		case "valor":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyValor(sn)
		}
		case "nexus":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyNexus(sn)
		}
		case "watsons":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsons(sn)
		}
		case "watsons_ha":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyZeratul(sn)
		}
	}
	tableData := [][]string{
		{sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	for _, v := range tableData {
		table.Append(v)
	}
	table.Render()
}

// SearchMany 6.x/7.x
func SearchMany(snMany []string) {
	var tableData = make([][]string, 0)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = threadQueryMode(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}
	fmt.Printf("cpemode: %s\n", mode)

	for _, sn := range snMany {
		tables := make([]string, 0)
		switch mode {
			case "valor":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyValor(sn)
			}
			case "nexus":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyNexus(sn)
			}
			case "watsons":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsons(sn)
			}
			case "watsons_ha":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyWatsonsHa(sn)
			}
			case "tassadar":{
				model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime = getCpebyZeratul(sn)
			}
		}
		tableData = append(tableData, append(tables, sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip))
	}
	// for _, v := range tableData {
	// 	table.Append(v)
	// }
	table.AppendBulk(tableData)
	table.Render()
}