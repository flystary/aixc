package net

import (
	"fmt"
	"os"
	"sort"

	. "aixc/tools/color"
	. "aixc/tools/table"
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

	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	if mode == "unknown" {
		os.Exit(13)
	}
	SyncDataMemorybyMode(mode)

	switch mode {
	case "valor":
		{
			ucpe = getCpebyValor(sn)
		}
	case "yifeng":
		{
			ucpe = getCpebyYifeng(sn)
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
	TableMarkdown(ucpes)
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
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	// 同步数据到内存
	SyncDataMemorybyMode(mode)
	for _, sn := range snMany {
		switch mode {
		case "valor":
			{
				ucpe = getCpebyValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "yifeng":
			{
				ucpe = getCpebyYifeng(sn)
				MaxVersion = cpeMaxVersionYifeng()
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
	TableBasic(ucpes)
}

// Search 6.x/7.x   单个Sn和随机mode
func SearchBySn(sn string) {
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	mode = ThreadQueryMode(sn)
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	if mode == "unknown" {
		os.Exit(14)
	}
	switch mode {
	case "valor":
		{
			ucpe = getCpebyValor(sn)

		}
	case "yifeng":
		{
			ucpe = getCpebyYifeng(sn)

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
	TableBasic(ucpes)
}

// SearchMany 6.x/7.x 多个Sn和随机mode
func SearchManyBySns(snMany []string) {
	msMap := make(map[string]string, len(snMany))
	// 同步所有数据
	syncCpe()
	// 多线程查询 属于哪个平台 并把对应数据存入内存
	for _, sn := range snMany {
		m := SearchSnByMode(sn)
		if m == "unknown" {
			continue
		}
		msMap[sn] = m
	}

	mdMap := make(map[string][][]string, 0)
	for _, sn := range snMany {
		mode = msMap[sn]
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "yifeng":
			{
				ucpe = uCPEInfoYifeng(sn)
				MaxVersion = cpeMaxVersionYifeng()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "watsons":
			{
				ucpe = uCPEInfoWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "watsonsha":
			{
				ucpe = uCPEInfoWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "tassadar":
			{
				ucpe = uCPEInfoZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		}

		if _, ok := mdMap[mode];!ok {
			// ucpes = append(ucpes, ucpe)
			ss := make([][]string, 0)
			ss = append(ss, ucpe)
			mdMap[mode] = ss
		} else {
			mdMap[mode] = append(mdMap[mode], ucpe)
		}
	}
	for mode, data := range mdMap {
		fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
		Table2Basic(data)
		fmt.Println()
	}
	// sort.Sort(ucpes.NotNull())
	// Table2Basic(ucpes)
}

// SearchByModeSns 所属mode和SnS
func SearchByModeSns(mode string, snMany []string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range snMany {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "yifeng":
			{
				ucpe = uCPEInfoYifeng(sn)
				MaxVersion = cpeMaxVersionYifeng()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "watsons":
			{
				ucpe = uCPEInfoWatsons(sn)
				MaxVersion = cpeMaxVersionWatsons()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "watsonsha":
			{
				ucpe = uCPEInfoWatsonsHa(sn)
				MaxVersion = cpeMaxVersionWatsonsHa()
				ucpe.NotNull().Version(MaxVersion).Time()
			}
		case "tassadar":
			{
				ucpe = uCPEInfoZeratul(sn)
				MaxVersion = cpeMaxVersionZeratul()
			}
		}
		ucpes = append(ucpes, ucpe)
	}
	// sort.Sort(ucpes.NotNull())
	Table2Basic(ucpes)
}

func SearchByMode(mode string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncDataMemorybyMode(mode)
	sn := "ALL"
	switch mode {
	case "valor":
		{
			MaxVersion = cpeMaxVersionValor()
			ucpes = allValor(sn)
		}
	case "yifeng":
		{
			MaxVersion = cpeMaxVersionYifeng()
			ucpes = allYifeng(sn)
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
	Table2Basic(ucpes)
}

// filter
// SearchByModeModel 所属mode和model
func FilterModelByMode(mode, model string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range getSnsByModel(mode, model) {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "yifeng":
			{
				ucpe = uCPEInfoYifeng(sn)
				MaxVersion = cpeMaxVersionYifeng()
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
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	Table2Basic(ucpes)
}

// FilterVersionByMode 所属mode和Version
func FilterVersionByMode(mode, version string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range getSnsByVersion(mode, version) {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "yifeng":
			{
				ucpe = uCPEInfoYifeng(sn)
				MaxVersion = cpeMaxVersionYifeng()
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
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes.NotNull())
	Table2Basic(ucpes)
}

// FilterPopByMode 所属mode和pop addr
func FilterPopByMode(mode, addr string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncDataMemorybyMode(mode)
	for _, sn := range getSnsByPopAddr(mode, addr) {
		switch mode {
		case "valor":
			{
				ucpe = uCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "yifeng":
			{
				ucpe = uCPEInfoYifeng(sn)
				MaxVersion = cpeMaxVersionYifeng()
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
		}
		ucpe.NotNull().Version(MaxVersion).Time()
		ucpes = append(ucpes, ucpe)
	}
	sort.Sort(ucpes.NotNull())
	Table2Basic(ucpes)
}

// FilterEnterpriseByMode 所属mode和企业号
func FilterEnterpriseByMode(mode, en string) {
	fmt.Printf("CPE %s is: %s\n", Blue("Mode"), White(mode))
	SyncEnDataMemorybyMode(mode)
	for _, sn := range getSnsByModeEn(mode, en) {
		switch mode {
		case "valor":
			{
				ucpe = EuCPEInfoValor(sn)
				MaxVersion = cpeMaxVersionValor()
			}
		case "yifeng":
			{
				ucpe = EuCPEInfoYifeng(sn)
				MaxVersion = cpeMaxVersionYifeng()
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
	Table3Basic(ucpes)
}
