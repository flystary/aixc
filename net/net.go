package net

import (
	"os"

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

// 加载url相关路由
func init() {
	path := "/etc/aixc/url.rules"
	if cr, err = loadURL(path); err != nil {
		os.Exit(9)
	}
	if token, err = getToken(cr.TokenRouteByMode()); err != nil {
		os.Exit(10)
	}
}

func snInSevenMode(sn string) string {
	if opo, err = getOperationData(token, cr.OperationRouteByMode()); err != nil {
		os.Exit(11)
	}
	return opo.SnInMode(sn)
}

func isSnInMode(sn, mode string) bool {
	URL := cr.CpeRouteByMode(mode)
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

func getCpePop(mode string) (string, string, string, string, string, string, string) {
	var model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime string
	popURL := cr.PopRouteByMode(mode)
	switch mode {
	case "valor":
		{
			model = vc.Model
			version = vc.SoftwareVersion
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
				masterpopip = mp.PopIP
			}
			if ok, bp = pv.IsId(backuppopid); ok != false {
				backuppopip = bp.PopIP
			}
		}
	case "nexus":
		{
			model = nb.Model
			version = nb.SoftwareVersion
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
				masterpopip = me.EntryIP
			}
			if ok, be = pn.IsId(backuppopid); ok != false {
				backuppopip = be.EntryIP
			}
		}
	case "watsons":
		{
			model = wv.Model
			version = wv.SoftwareVersion
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
				masterpopip = me.EntryIP
			}
			if ok, be = pw.IsId(backuppopid); ok != false {
				backuppopip = be.EntryIP
			}
		}
	case "watsons_ha":
		{
			model = whv.Model
			version = whv.SoftwareVersion
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
				masterpopip = me.EntryIP
			}
			if ok, be = pwh.IsId(backuppopid); ok != false {
				backuppopip = be.EntryIP
			}
		}
	case "tassadar":
		{
			model = zc.Model
			version = zc.SoftwareVersion
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
				masterpopip = ms.EntryIP
			}
			if ok, bs = pz.IsId(backuppopid); ok != false {
				backuppopip = bs.EntryIP
			}
		}
	}
	return model, version, mastercpeip, masterpopip, backupcpeip, backuppopip, updatetime
}

func getDvc(sn, mode string) {
	dvcURL := cr.DeviceRouteByMode(mode)
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
			ok, zd = d.IsSn(sn)
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
			ok, whd = d.IsSn(sn)
			if ok == false {
				os.Exit(16)
			}
		}
	}
}

// seven
//传入sn
//根据sn获取盒子版本
//

// six
