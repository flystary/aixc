package net


import (
	"os"
	"fmt"
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
	mode = getModebySevenSn(sn)
	fmt.Printf("sevencpemode: %s\n", mode)
	if mode == "unknown" {
		os.Exit(13)
	}
	syncDataMemorybyMode(mode)
	switch mode {
		case "valor":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyValor(sn)
		}
		case "nexus":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyNexus(sn)
		}
		case "watsons":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsons(sn)
		}
		case "watsons_ha":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyZeratul(sn)
		}
	}
	tableData := [][]string{
		{sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip},
	}
	tableMarkdown(tableData)
}

// SearchSevenMany 6.x
func SearchSevenMany(snMany []string) {
	// table
	var tableData = make([][]string, 0)
	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = getModebySevenSn(sn)
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
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyValor(sn)
			}
			case "nexus":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyNexus(sn)
			}
			case "watsons":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsons(sn)
			}
			case "watsons_ha":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsonsHa(sn)
			}
			case "tassadar":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyZeratul(sn)
			}
		}
		tableData = append(tableData, append(tables, sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip))
	}
	tableBasic(tableData)
}

//Search 6.x/7.x
func Search(sn string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	mode = threadQueryMode(sn)
	fmt.Printf("cpemode: %s\n", mode)
	if mode == "unknown" {
		os.Exit(14)
	}
	switch mode {
		case "valor":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyValor(sn)
		}
		case "nexus":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyNexus(sn)
		}
		case "watsons":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsons(sn)
		}
		case "watsons_ha":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyZeratul(sn)
		}
	}
	tableData := [][]string{
		{sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip},
	}
	tableMarkdown(tableData)
}

// SearchMany 6.x/7.x
func SearchMany(snMany []string) {
	var tableData = make([][]string, 0)
	// 多线程查询 属于哪个平台 并把对应数据存入内存
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
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyValor(sn)
			}
			case "nexus":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyNexus(sn)
			}
			case "watsons":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsons(sn)
			}
			case "watsons_ha":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyWatsonsHa(sn)
			}
			case "tassadar":{
				model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip = getCpebyZeratul(sn)
			}
		}
		tableData = append(tableData, append(tables, sn, model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip))
	}
	tableBasic(tableData)
}