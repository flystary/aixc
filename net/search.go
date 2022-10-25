package net

import (
	"fmt"
	"os"
	"strings"
	"time"
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
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
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
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
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

		for i, length := 0, len(ucpes); i < length; i++ {
			if ucpes[i] == "" {
				break
			}
			//版本不一致 区别显示
			if i == 2 {
				if VERSION == "" {
					VERSION  = ucpes[2]
					ucpes[2] = Cyan(ucpes[2])
				} else {
					if ucpes[2] != VERSION {
						ucpes[2] = Red(ucpes[2])
					}
				}
			}
			// 入口同步时间不是今天1小时内 区别显示
			if i == 3 {
				var now = time.Now()
				synctime, _ := time.Parse("2006-01-02 15:04:05", ucpes[3])
				if synctime.Year() != now.Year() || synctime.Month() != now.Month() || synctime.Day() != now.Day() || synctime.Hour() != now.Hour() {
					ucpes[3] = fmt.Sprintf("%s✗%s", Red(strings.Split(ucpes[3], " ")[0]), Red(strings.Split(ucpes[3], " ")[1]))
				}
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
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
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
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))

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

		for i, length := 0, len(ucpes); i < length; i++ {
			if ucpes[i] == "" {
				break
			}
			//版本不一致 区别显示
			if i == 2 {
				if VERSION == "" {
					VERSION  = ucpes[2]
					ucpes[2] = Cyan(ucpes[2])
				} else {
					if ucpes[2] != VERSION {
						ucpes[2] = Red(ucpes[2])
					}
				}
			}
			// 入口同步时间不是今天1小时内 区别显示
			if i == 3 {
				var now = time.Now()
				synctime, _ := time.Parse("2006-01-02 15:04:05", ucpes[3])
				if synctime.Year() != now.Year() || synctime.Month() != now.Month() || synctime.Day() != now.Day() || synctime.Hour() != now.Hour() {
					ucpes[3] = fmt.Sprintf("%s✗%s", Red(strings.Split(ucpes[3], " ")[0]), Red(strings.Split(ucpes[3], " ")[1]))
				}
			}
		}
		slices = append(slices, ucpes)
	}
	tableBasic(slices)
}