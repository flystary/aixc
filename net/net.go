package net

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"aixc/conf"
	"aixc/struct/cpe"
	"aixc/struct/dvc"
	"aixc/struct/opt"
	"aixc/struct/pop"
)

var (
	cr    conf.Rules
	token string
	err   error
	ok    bool

	zc  cpe.Cpe  //zeratul
	vc  cpe.Cpe  //valor
	nb  cpe.Box  //nexus
	wv  cpe.VBox //watsons
	whv cpe.VBox //watsonsHa

	nd  dvc.Ndvc //nexus
	vd  dvc.Vdvc //valor
	zd  dvc.Vdvc //zeratul
	wd  dvc.Wdvc //watsons
	whd dvc.Wdvc //watsonsha

	mp pop.Pop   //masterpop
	bp pop.Pop   //backuppop
	me pop.Entry //masterentry
	be pop.Entry //backupentry
	ms pop.SPop  //masterspop
	bs pop.SPop  //backupspop

	opo opt.Operation
)

func init() {
	path := "./url.rules"
	if cr, err = initURL(path); err != nil {
		os.Exit(9)
	}
	if token, err = getResponseToken(cr.GetTokenUrlPath()); err != nil {
		os.Exit(10)
	}
}

//SearchSeven 6.x
func SearchSeven(sn string) {
	// 已知mode数据6.x版本
	m := snInSevenMode(sn)
	fmt.Printf("sevencpemode: %s\n", m)
	if m == "unknown" {
		os.Exit(13)
	}
	isSnInMode(sn, m)
	mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime := getCpePop(m)
	fmt.Printf("updatetime: %s", updatetime)
	fmt.Printf("masterpop: %s mastercpe: %s", masterpopip, mastercpeip)
	fmt.Printf("\nbackuppop: %s backupcpe: %s\n", backuppopip, backupcpeip)

	getDvc(sn, m)
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
	mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime := getCpePop(m)
	tableData := [][]string{
		{sn, vc.Model, vc.SoftwareVersion, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip},
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

	// 多线程查询 属于哪个平台
	for _, sn := range snMany {
		mode = getSnInMode(sn)
		if mode == "unknown" {
			os.Exit(14)
		}
		break
	}

	for _, sn := range snMany {
		tables := make([]string, 0)
		isSnInMode(sn, mode)
		mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime := getCpePop(mode)
		tableData = append(tableData, append(tables, sn, vc.Model, vc.SoftwareVersion, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip))
	}

	for _, v := range tableData {
		table.Append(v)
	}
	table.Render()
}

func snInSevenMode(sn string) string {
	if opo, err = getOperationData(token, cr.GetOperationUrlPath()); err != nil {
		os.Exit(11)
	}
	return opo.SnInMode(sn)
}

func isSnInMode(sn, mode string) bool {
	URL := cr.GetCpeUrlPath(mode)
	switch mode {
	case "valor":
		{
			cv, err := getValorData(token, URL)
			if err != nil {
				os.Exit(12)
			}
			if ok, vc = cv.IsSn(sn); ok != false {
				return true
			}
		}
	case "nexus":
		{
			cn, err := getNexusData(token, URL)
			if err != nil {
				os.Exit(12)
			}
			if ok, nb = cn.IsSn(sn); ok != false {
				return true
			}
		}
	case "watsons":
		{
			cw, err := getWatsonsData(token, URL)
			if err != nil {
				os.Exit(12)
			}
			if ok, wv = cw.IsSn(sn); ok != false {
				return true
			}
		}
	case "watsons_ha":
		{
			cwh, err := getWatsonsHaData(token, URL)
			if err != nil {
				os.Exit(12)
			}
			if ok, whv = cwh.IsSn(sn); ok != false {
				return true
			}
		}
	case "tassadar":
		{
			cz, err := getZeratulData(token, URL)
			if err != nil {
				os.Exit(12)
			}
			if ok, zc = cz.IsSn(sn); ok != false {
				return true
			}
		}
	}
	return false
}

func getCpePop(mode string) (string, string, string, string, string) {
	var mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime string
	popURL := cr.GetPopUrlPath(mode)
	switch mode {
	case "valor":
		{
			mastercpeip = vc.MasterPopIP
			backupcpeip = vc.BackupPopIP
			masterpopid := vc.MasterPopID
			backuppopid := vc.BackupPopID
			updatetime = vc.EntryUpdateTime
			pv, err := getValorPopData(token, popURL)
			if err != nil {
				os.Exit(12)
			}
			if ok, mp = pv.IsId(masterpopid); ok != false {
				masterpopip = pv.GetPopStruct(masterpopid).PopIP
			}
			if ok, bp = pv.IsId(backuppopid); ok != false {
				backuppopip = pv.GetPopStruct(backuppopid).PopIP
			}
		}
	case "nexus":
		{
			mastercpeip = nb.MasterEntryIP
			backupcpeip = nb.BackupEntryIP
			masterpopid := nb.MasterEntryID
			backuppopid := nb.BackupEntryID
			updatetime = nb.EntryUpdateTime
			pn, err := getNexusEntryData(token, popURL)
			if err != nil {
				os.Exit(12)
			}
			if ok, me = pn.IsId(masterpopid); ok != false {
				masterpopip = pn.GetPopStruct(masterpopid).EntryIP
			}
			if ok, be = pn.IsId(backuppopid); ok != false {
				backuppopip = pn.GetPopStruct(backuppopid).EntryIP
			}
		}
	case "watsons":
		{
			mastercpeip = wv.MasterEntryIP
			backupcpeip = wv.BackupEntryIP
			masterpopid := wv.MasterEntryID
			backuppopid := wv.BackupEntryID
			updatetime = wv.EntryUpdateTime
			pw, err := getWatsonsEntryData(token, popURL)
			if err != nil {
				os.Exit(12)
			}
			if ok, me = pw.IsId(masterpopid); ok != false {
				masterpopip = pw.GetPopStruct(masterpopid).EntryIP
			}
			if ok, be = pw.IsId(backuppopid); ok != false {
				backuppopip = pw.GetPopStruct(backuppopid).EntryIP
			}
		}
	case "watsons_ha":
		{
			mastercpeip = whv.MasterEntryIP
			backupcpeip = whv.BackupEntryIP
			masterpopid := whv.MasterEntryID
			backuppopid := whv.BackupEntryID
			updatetime = whv.EntryUpdateTime
			pwh, err := getWatsonsHaEntryData(token, popURL)
			if err != nil {
				os.Exit(12)
			}
			if ok, me = pwh.IsId(masterpopid); ok != false {
				masterpopip = pwh.GetPopStruct(masterpopid).EntryIP
			}
			if ok, be = pwh.IsId(backuppopid); ok != false {
				backuppopip = pwh.GetPopStruct(backuppopid).EntryIP
			}
		}
	case "tassadar":
		{
			mastercpeip = zc.MasterPopIP
			backupcpeip = zc.BackupPopIP
			masterpopid := zc.MasterPopID
			backuppopid := zc.BackupPopID
			updatetime = zc.EntryUpdateTime
			pz, err := getZeratulPopData(token, popURL)
			if err != nil {
				os.Exit(12)
			}
			if ok, ms = pz.IsId(masterpopid); ok != false {
				masterpopip = pz.GetPopStruct(masterpopid).EntryIP
			}
			if ok, bs = pz.IsId(backuppopid); ok != false {
				backuppopip = pz.GetPopStruct(backuppopid).EntryIP
			}
		}
	}
	return mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime
}

func getDvc(sn, mode string) {
	dvcURL := cr.GetDvcUrlPath(mode)
	switch mode {
	case "nexus":
		{
			d, err := getDvcNexusData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			ok, nd = d.IsSn(sn)
			if ok == false {
				os.Exit(16)
			}
		}
	case "valor":
		{
			d, err := getDvcValorData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			ok, vd = d.IsSn(sn)
			if ok == false {
				os.Exit(16)
			}
		}
	case "tassadar":
		{
			d, err := getDvcZeratulData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			ok, vd = d.IsSn(sn)
			if ok == false {
				os.Exit(16)
			}
		}
	case "watsons":
		{
			d, err := getDvcWatsonsData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			ok, wd = d.IsSn(sn)
			if ok == false {
				os.Exit(16)
			}
		}
	case "watsons_ha":
		{
			d, err := getDvcWatsonsHaData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			ok, wd = d.IsSn(sn)
			if ok == false {
				os.Exit(16)
			}
		}
	}
}
