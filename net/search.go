package net

import (
	co "aixc/utils/color"
	tb "aixc/utils/table"
	"fmt"
	"os"
	"sort"
)

var (
	mode       string
	MaxVersion string // VERSION 当前版本的最大
	ucpes      Ucpes  = make([][]string, 0)
	ucpe       Ucpe   = make([]string, 0)
)

// SearchSeven      单个Sn属于6.x
func SearchSevenBySn(sn string) {
	mode = GetModebySevenSn(sn)

	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	if mode == "unknown" {
		os.Exit(13)
	}
	SyncDataMemorybyMode(mode)

	switch mode {
	case "valor":
		{
			ucpe = getCpebyValor(sn)
		}
	case "nexus":
		{
			ucpe = getCpebyNexus(sn)
		}
	case "watsons":
		{
			ucpe = getCpebyWatsons(sn)
		}
	case "watsonsha":
		{
			ucpe = getCpebyWatsonsHa(sn)
		}
	case "tassadar":
		{
			ucpe = getCpebyZeratul(sn)
		}
	}
	ucpes = append(ucpes, ucpe)
	tb.TableMarkdown(ucpes)
}

// SearchSevenMany 多个Sn属于6.x
func SearchSevenManyBySns(snMany []string) {
	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = GetModebySevenSn(sn)
		if mode == "unknown" {
			os.Exit(14)
		} else {
			break
		}

	}
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	// 同步数据到内存
	SyncDataMemorybyMode(mode)
	for _, sn := range snMany {
		switch mode {
		case "valor":
			{
				ucpe = getCpebyValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "nexus":
			{
				ucpe = getCpebyNexus(sn)
				MaxVersion = cpeMaxVersionNexus()
			}
		case "watsons":
			{
				ucpe = getCpebyWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
			}
		case "watsonsha":
			{
				ucpe = getCpebyWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
			}
		case "tassadar":
			{
				ucpe = getCpebyZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes)
	tb.TableBasic(ucpes)
}

// Search 6.x/7.x   单个Sn和随机mode
func SearchBySn(sn string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	mode = ThreadQueryMode(sn)
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	if mode == "unknown" {
		os.Exit(14)
	}
	switch mode {
	case "valor":
		{
			ucpe = getCpebyValor(sn)

		}
	case "nexus":
		{
			ucpe = getCpebyNexus(sn)
		}
	case "watsons":
		{
			ucpe = getCpebyWatsons(sn)
		}
	case "watsonsha":
		{
			ucpe = getCpebyWatsonsHa(sn)
		}
	case "tassadar":
		{
			ucpe = getCpebyZeratul(sn)
		}
	}
	ucpes = append(ucpes, ucpe)
	tb.TableBasic(ucpes)
}

// SearchMany 6.x/7.x 多个Sn和随机mode
func SearchManyBySns(snMany []string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	for _, sn := range snMany {
		mode = ThreadQueryMode(sn)
		if mode == "unknown" {
			os.Exit(14)
		} else {
			break
		}
	}

	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))

	for _, sn := range snMany {
		switch mode {
		case "valor":
			{
				ucpe = getCpebyValor(sn)
				MaxVersion = cpeMaxVersionValor()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "nexus":
			{
				ucpe = getCpebyNexus(sn)
				MaxVersion = cpeMaxVersionNexus()
			}
		case "watsons":
			{
				ucpe = getCpebyWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "watsonsha":
			{
				ucpe = getCpebyWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "tassadar":
			{
				ucpe = getCpebyZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		}
		ucpes = append(ucpes, ucpe)
	}
	// sort.Sort(ucpes)
	tb.TableBasic(ucpes)
}

// SearchByModeSns 所属mode和SnS
func SearchByModeSns(mode string, sns []string) {
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range sns {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "nexus":
			{
				ucpe = uCPEInfoNexus(sn)
				MaxVersion = cpeMaxVersionNexus()
			}
		case "watsons":
			{
				ucpe = uCPEInfoWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
			}
		case "watsonsha":
			{
				ucpe = uCPEInfoWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
			}
		case "tassadar":
			{
				ucpe = uCPEInfoZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	// sort.Sort(ucpes.NotNull())
	tb.Table2Basic(ucpes)
}

func SearchByMode(mode string) {
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	SyncDataMemorybyMode(mode)
	sn := "ALL"
	switch mode {
	case "valor":
		{
			MaxVersion = cpeMaxVersionValor()
			ucpes = allValor(sn)
		}
	case "tassadar":
		{
			MaxVersion = cpeMaxVersionZeratul()
			ucpes = allZeratul(sn)
		}
	case "watsonsha":
		{
			MaxVersion = cpeMaxVersionWatsonsHa()
			ucpes = allWatsonsHa(sn)
		}
	}
	sort.Sort(ucpes)
	tb.Table2Basic(ucpes)
}

// filter
// SearchByModeModel 所属mode和model
func FilterModelByMode(mode, model string) {
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range getSnsByModel(mode, model) {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "tassadar":
			{
				ucpe = uCPEInfoZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		case "watsons":
			{
				ucpe = uCPEInfoWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
			}
		case "watsonsha":
			{
				ucpe = uCPEInfoWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
			}
		case "nexus":
			{
				ucpe = uCPEInfoNexus(sn)
				MaxVersion = cpeMaxVersionNexus()
			}
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	tb.Table2Basic(ucpes)
}

// FilterVersionByMode 所属mode和Version
func FilterVersionByMode(mode, version string) {
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range getSnsByVersion(mode, version) {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "tassadar":
			{
				ucpe = uCPEInfoZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		case "watsons":
			{
				ucpe = uCPEInfoWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
			}
		case "watsonsha":
			{
				ucpe = uCPEInfoWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
			}
		case "nexus":
			{
				ucpe = uCPEInfoNexus(sn)
				MaxVersion = cpeMaxVersionNexus()
			}

		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes.NotNull())
	tb.Table2Basic(ucpes)
}

// FilterPopByMode 所属mode和pop addr
func FilterPopByMode(mode, addr string) {
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range getSnsByPopAddr(mode, addr) {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "tassadar":
			{
				ucpe = uCPEInfoZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		case "watsons":
			{
				ucpe = uCPEInfoWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
			}
		case "watsonsha":
			{
				ucpe = uCPEInfoWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
			}
		case "nexus":
			{
				ucpe = uCPEInfoNexus(sn)
				MaxVersion = cpeMaxVersionNexus()
			}
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes.NotNull())
	tb.Table2Basic(ucpes)
}

// FilterEnterpriseByMode 所属mode和企业号
func FilterEnterpriseByMode(mode, en string) {
	fmt.Printf("CPE %s is: %s\n", co.Blue("Mode"), co.White(mode))
	SyncEnDataMemorybyMode(mode)
	for _, sn := range getSnsByModeEn(mode, en) {
		switch mode {
		case "valor":
			{
				ucpe = EuCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "nexus":
			{
				ucpe = EuCPEInfoNexus(sn)
				MaxVersion = cpeMaxVersionNexus()
			}
		case "tassadar":
			{
				ucpe = EuCPEInfoZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes.NotNull())
	tb.Table3Basic(ucpes)
}
