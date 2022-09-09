package net

import (
	"aixc/struct/cpe"
	"aixc/struct/dvc"
	"aixc/struct/pop"
	"os"
)


var (
	cv  cpe.Valor     //valor
	cn  cpe.Nexus     //nexus
	cw  cpe.Watsons   //watsons
	ch  cpe.WatsonsHa //watsonsha
	cz 	cpe.Zeratul   //zeratul

	pv 	pop.Valor
	pn	pop.Nexus
	pw	pop.Watsons
	ph 	pop.WatsonsHa
	pz	pop.Zeratul

	dv  dvc.Valor
	dn  dvc.Nexus
	dw  dvc.Watsons
	dh  dvc.WatsonsHa
	dz  dvc.Zeratul
)

type ucpe struct {
	model   	string
	version 	string
	mastercpeip string
	masterpopip string
	backupcpeip string
	backuppopip string
	updatetime 	string
}

// 已知mode获取cpe,dvc,pop数据并放入到内存
func syncDataMemorybyMode(mode string) {
	cpeURL := rules.CpeRouteByMode(mode)
	popURL := rules.PopRouteByMode(mode)
	dvcURL := rules.DeviceRouteByMode(mode)

	switch mode {
		case "valor":{
			cv, err = getValorData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pv, err = getValorPopData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			dv, err = getDvcValorData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
		}
		case "nexus":{
			cn, err = getNexusData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pn, err = getNexusEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			dn, err = getDvcNexusData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
		}
		case "watsons":{
			cw, err = getWatsonsData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pw, err = getWatsonsEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			dw, err = getDvcWatsonsData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
		}
		case "watsons_ha":{
			ch, err = getWatsonsHaData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			ph, err = getWatsonsHaEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			dh, err = getDvcWatsonsHaData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
		}
		case "tassadar":{
			cz, err = getZeratulData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pz, err = getZeratulPopData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			dz, err = getDvcZeratulData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
		}
	}
}

// 根据sn和mode获取cpe,dvc,pop数据并放入到内存
func syncDataMemorybySnMode(sn, mode string) bool {
	cpeURL := rules.CpeRouteByMode(mode)
	popURL := rules.PopRouteByMode(mode)
	dvcURL := rules.DeviceRouteByMode(mode)
	switch mode {
		case "valor":
			{
				cv, err = getValorData(token, cpeURL)
				if err != nil {
					os.Exit(12)
				}
				pv, err = getValorPopData(token, popURL)
				if err != nil {
					os.Exit(13)
				}
				dv, err = getDvcValorData(token, dvcURL)
				if err != nil {
					os.Exit(15)
				}
				if ok, _ := cv.IsSn(sn); ok != false {
					return true
				}
			}
		case "nexus":
			{
				cn, err = getNexusData(token, cpeURL)
				if err != nil {
					os.Exit(12)
				}
				pn, err = getNexusEntryData(token, popURL)
				if err != nil {
					os.Exit(13)
				}
				dn, err = getDvcNexusData(token, dvcURL)
				if err != nil {
					os.Exit(15)
				}
				if ok, _ := cn.IsSn(sn); ok != false {
					return true
				}
			}
		case "watsons":
			{
				cw, err = getWatsonsData(token, cpeURL)
				if err != nil {
					os.Exit(12)
				}
				pw, err = getWatsonsEntryData(token, popURL)
				if err != nil {
					os.Exit(13)
				}
				dw, err = getDvcWatsonsData(token, dvcURL)
				if err != nil {
					os.Exit(15)
				}
				if ok, _ := cw.IsSn(sn); ok != false {
					return true
				}
			}
		case "watsons_ha":
			{
				ch, err = getWatsonsHaData(token, cpeURL)
				if err != nil {
					os.Exit(12)
				}
				ph, err = getWatsonsHaEntryData(token, popURL)
				if err != nil {
					os.Exit(13)
				}
				dh, err = getDvcWatsonsHaData(token, dvcURL)
				if err != nil {
					os.Exit(15)
				}
				if ok, _ := ch.IsSn(sn); ok != false {
					return true
				}
			}
		case "tassadar":
			{
				cz, err = getZeratulData(token, cpeURL)
				if err != nil {
					os.Exit(12)
				}
				pz, err = getZeratulPopData(token, popURL)
				if err != nil {
					os.Exit(13)
				}
				dz, err = getDvcZeratulData(token, dvcURL)
				if err != nil {
					os.Exit(15)
				}
				if ok, _ := cz.IsSn(sn); ok != false {
					return true
				}
			}
		}
	return false
}

func getCpebyValor(sn string) (string, string, string, string, string, string, string) {
	cpe := cv.GetCpeStructBySn(sn)

	model   	:= cpe.Model
	version 	:= cpe.SoftwareVersion
	updatetime	:= cpe.EntryUpdateTime
	mastercpeip := cpe.MasterPopIP
	masterpopip := pv.GetPopStructById(cpe.MasterPopID).PopIP
	backupcpeip := cpe.BackupPopIP
	backuppopip := pv.GetPopStructById(cpe.BackupPopID).PopIP
	
	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyNexus(sn string) (string, string, string, string, string, string, string) {
	cpe := cn.GetCpeStructBySn(sn)

	model   	:= cpe.Model
	version 	:= cpe.SoftwareVersion
	updatetime	:= cpe.EntryUpdateTime
	mastercpeip := cpe.MasterEntryIP
	masterpopip := pn.GetPopStructById(cpe.MasterEntryID).EntryIP
	backupcpeip := cpe.BackupEntryIP
	backuppopip := pn.GetPopStructById(cpe.BackupEntryID).EntryIP
	
	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyWatsons(sn string) (string, string, string, string, string, string, string) {
	cpe := cw.GetCpeStructBySn(sn)
	
	model   	:= cpe.Model
	version 	:= cpe.SoftwareVersion
	updatetime	:= cpe.EntryUpdateTime
	mastercpeip := cpe.MasterEntryIP
	masterpopip := pw.GetPopStructById(cpe.MasterEntryID).EntryIP
	backupcpeip := cpe.BackupEntryIP
	backuppopip := pw.GetPopStructById(cpe.BackupEntryID).EntryIP
	
	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyWatsonsHa(sn string) (string, string, string, string, string, string, string) {
	cpe := ch.GetCpeStructBySn(sn)

	model   	:= cpe.Model
	version 	:= cpe.SoftwareVersion
	updatetime	:= cpe.EntryUpdateTime
	mastercpeip := cpe.MasterEntryIP
	masterpopip := ph.GetPopStructById(cpe.MasterEntryID).EntryIP
	backupcpeip := cpe.BackupEntryIP
	backuppopip := ph.GetPopStructById(cpe.BackupEntryID).EntryIP

	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyZeratul(sn string) (string, string, string, string, string, string, string) {
	cpe := cz.GetCpeStructBySn(sn)

	model   	:= cpe.Model
	version 	:= cpe.SoftwareVersion
	updatetime	:= cpe.EntryUpdateTime
	mastercpeip := cpe.MasterPopIP
	masterpopip := pz.GetPopStructById(cpe.MasterPopID).EntryIP
	backupcpeip := cpe.BackupPopIP
	backuppopip := pz.GetPopStructById(cpe.BackupPopID).EntryIP
	
   return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}