package net

import (
	"fmt"
	"os"
	"sort"
)

var (
	mode	 string
	// VERSION 当前版本的最大
	MaxVersion string
	ucpes Ucpes = make([][]string, 0)
	ucpe Ucpe = make([]string, 0)
	// ucpe  = make([]string, 0)
)

//SearchSeven      单个Sn属于6.x
func SearchSevenBySn(sn string) {
	mode = GetModebySevenSn(sn)
	// red := color.New(color.FgBlue, color.Bold).SprintFunc()
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	if mode == "unknown" {
		os.Exit(13)
	}
	SyncDataMemorybyMode(mode)

	switch mode {
		case "valor":{
			ucpe = GetCpebyValor(sn)
		}
		case "nexus":{
			ucpe = GetCpebyNexus(sn)
		}
		case "watsons":{
			ucpe = GetCpebyWatsons(sn)
		}
		case "watsonsha":{
			ucpe = GetCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			ucpe = GetCpebyZeratul(sn)
		}
	}
	ucpes = append(ucpes, ucpe)
	tableMarkdown(ucpes)
}

// SearchSevenMany 多个Sn属于6.x
func SearchSevenManyBySns(snMany []string) {
	// table
	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = GetModebySevenSn(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	// 同步数据到内存
	SyncDataMemorybyMode(mode)
	for _, sn := range snMany {
		switch mode {
			case "valor":{
				ucpe = GetCpebyValor(sn)
				MaxVersion = GetCpeMaxVsbyValor()
			}
			case "nexus":{
				ucpe = GetCpebyNexus(sn)
				MaxVersion = GetCpeMaxVsbyNexus()
			}
			case "watsons":{
				ucpe = GetCpebyWatsons(sn)
				MaxVersion = GetCpeMaxVsbyWatsons()
			}
			case "watsonsha":{
				ucpe = GetCpebyWatsonsHa(sn)
				MaxVersion = GetCpeMaxVsbyWatsonsHa()
			}
			case "tassadar":{
				ucpe = GetCpebyZeratul(sn)
				MaxVersion = GetCpeMaxVsbyZeratul()
			}
		}
		ucpe.Null().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes)
	tableBasic(ucpes)
}

//Search 6.x/7.x   单个Sn和随机mode
func SearchBySn(sn string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	mode = ThreadQueryMode(sn)
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	if mode == "unknown" {
		os.Exit(14)
	}
	switch mode {
		case "valor":{
			ucpe = GetCpebyValor(sn)

		}
		case "nexus":{
			ucpe = GetCpebyNexus(sn)
		}
		case "watsons":{
			ucpe = GetCpebyWatsons(sn)
		}
		case "watsonsha":{
			ucpe = GetCpebyWatsonsHa(sn)
		}
		case "tassadar":{
			ucpe = GetCpebyZeratul(sn)
		}
	}
	ucpes = append(ucpes, ucpe)
	tableBasic(ucpes)
}

// SearchMany 6.x/7.x 多个Sn和随机mode
func SearchManyBySns(snMany []string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	for _, sn := range snMany {
		mode = ThreadQueryMode(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}

	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))

	for _, sn := range snMany {
		switch mode {
			case "valor":{
				ucpe = GetCpebyValor(sn)
				MaxVersion = GetCpeMaxVsbyValor()
			}
			case "nexus":{
				ucpe = GetCpebyNexus(sn)
				MaxVersion = GetCpeMaxVsbyNexus()
			}
			case "watsons":{
				ucpe = GetCpebyWatsons(sn)
				MaxVersion = GetCpeMaxVsbyWatsons()
			}
			case "watsonsha":{
				ucpe = GetCpebyWatsonsHa(sn)
				MaxVersion = GetCpeMaxVsbyWatsonsHa()
			}
			case "tassadar":{
				ucpe = GetCpebyZeratul(sn)
				MaxVersion = GetCpeMaxVsbyZeratul()
			}
		}
		ucpe.Null().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes)
	tableBasic(ucpes)
}

//SearchByModeSns 所属mode和SnS
func SearchByModeSns(mode string,sns []string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range sns {
		switch mode {
			case "valor":{
				ucpe = GetCpebyValor(sn)
				MaxVersion = GetCpeMaxVsbyValor()
			}
			case "nexus":{
				ucpe = GetCpebyNexus(sn)
				MaxVersion = GetCpeMaxVsbyNexus()
			}
			case "watsons":{
				ucpe = GetCpebyWatsons(sn)
				MaxVersion = GetCpeMaxVsbyWatsons()
			}
			case "watsonsha":{
				ucpe = GetCpebyWatsonsHa(sn)
				MaxVersion = GetCpeMaxVsbyWatsonsHa()
			}
			case "tassadar":{
				ucpe = GetCpebyZeratul(sn)
				MaxVersion = GetCpeMaxVsbyZeratul()
			}
		}
		ucpe.Null().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes.Null())
	tableBasic(ucpes)
}

//SearchByEnterprise 所属mode和企业号
func SearchByModeEnterprise(mode,en string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncEnDataMemorybyMode(mode)
	for _, sn := range getSnsByMode(mode, en) {
		switch mode {
			case "valor":{
				ucpe = GetCpebyValor(sn)
				MaxVersion = GetCpeMaxVsbyValor()
			}
			case "nexus":{
				ucpe = GetCpebyNexus(sn)
				MaxVersion = GetCpeMaxVsbyNexus()
			}
			case "tassadar":{
				ucpe = GetCpebyZeratul(sn)
				MaxVersion = GetCpeMaxVsbyZeratul()
			}
		}
		ucpe.Null().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes.Null())
	tableBasic(ucpes)
}
