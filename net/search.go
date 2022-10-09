package net


import (
	"os"
	"fmt"
)

var (
	mode	string
	// VERSION 默认第一个版本
	VERSION string
	slices = make([][]string, 0)
	ucpes  = make([]string, 0)

)

//SearchSeven 6.x
func SearchSeven(sn string) {
	mode = getModebySevenSn(sn)
	// red := color.New(color.FgBlue, color.Bold).SprintFunc()
	fmt.Printf("CPE %s is: %s\n", blue("Mode"), white(mode))
	if mode == "unknown" {
		os.Exit(13)
	}
	syncDataMemorybyMode(mode)

	switch mode {
		case "valor":{
			ucpes = getCpebyValor(sn)
		}
		case "nexus":{
			ucpes = getCpebyNexus(sn)
		}
		case "watsons":{
			ucpes = getCpebyWatsons(sn)
		}
		case "watsonsha":{
			ucpes = getCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			ucpes = getCpebyZeratul(sn)
		}
	}
	slices = append(slices, ucpes)
	tableMarkdown(slices)
}

// SearchSevenMany 6.x
func SearchSevenMany(snMany []string) {
	// table
	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = getModebySevenSn(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}
	fmt.Printf("CPE %s is: %s\n", blue("Mode"), white(mode))
	// 同步数据到内存
	syncDataMemorybyMode(mode)
	for _, sn := range snMany {
		switch mode {
			case "valor":{
				ucpes = getCpebyValor(sn)
			}
			case "nexus":{
				ucpes = getCpebyNexus(sn)
			}
			case "watsons":{
				ucpes = getCpebyWatsons(sn)
			}
			case "watsonsha":{
				ucpes = getCpebyWatsonsHa(sn)
			}
			case "tassadar":{
				ucpes = getCpebyZeratul(sn)
			}
		}
		//版本号不一样区别显示
		if VERSION == ""{
			VERSION = ucpes[2]
		}else {
			if ucpes[2] != VERSION {
				ucpes[2] = red(ucpes[2])
			}
		}
		slices = append(slices, ucpes)
	}
	tableBasic(slices)
}

//Search 6.x/7.x
func Search(sn string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	mode = threadQueryMode(sn)
	fmt.Printf("CPE %s is: %s\n", blue("Mode"), white(mode))
	if mode == "unknown" {
		os.Exit(14)
	}
	switch mode {
		case "valor":{
			ucpes = getCpebyValor(sn)
		}
		case "nexus":{
			ucpes = getCpebyNexus(sn)
		}
		case "watsons":{
			ucpes = getCpebyWatsons(sn)
		}
		case "watsonsha":{
			ucpes = getCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			ucpes = getCpebyZeratul(sn)
		}
	}
	slices = append(slices, ucpes)
	tableMarkdown(slices)
}

// SearchMany 6.x/7.x
func SearchMany(snMany []string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	for _, sn := range snMany {
		mode = threadQueryMode(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}
	fmt.Printf("CPE %s is: %s\n", blue("Mode"), white(mode))

	for _, sn := range snMany {
		switch mode {
			case "valor":{
				ucpes = getCpebyValor(sn)
			}
			case "nexus":{
				ucpes = getCpebyNexus(sn)
			}
			case "watsons":{
				ucpes = getCpebyWatsons(sn)
			}
			case "watsonsha":{
				ucpes = getCpebyWatsonsHa(sn)
			}
			case "tassadar":{
				ucpes = getCpebyZeratul(sn)
			}
		}
		if VERSION == ""{
			VERSION = ucpes[2]
		}else {
			if ucpes[2] != VERSION {
				ucpes[2] = red(ucpes[2])
			}
		}

		slices = append(slices, ucpes)
	}
	tableBasic(slices)
}