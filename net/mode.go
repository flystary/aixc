package net

type ucpe struct {
	model       string
	version     string
	mastercpeip string
	masterpopip string
	backupcpeip string
	backuppopip string
	updatetime  string
}

//Mtype 自定义int
type Mtype int

const (
	valor Mtype = iota + 1
	nexus
	watsons
	watsonsHa
	tassadar
)

func (m Mtype) enum() string {
	var mode string
	switch m {
	case valor:
		{
			mode = "valor"
		}
	case nexus:
		{
			mode = "nexus"
		}
	case watsons:
		{
			mode = "watsons"
		}
	case watsonsHa:
		{
			mode = "watsons_ha"
		}
	case tassadar:
		{
			mode = "tassadar"
		}
	default:
		{
			mode = "unknown"
		}
	}
	return mode
}

func getCpebyValor(sn string) (string, string, string, string, string, string, string) {
	cpe := cv.GetCpeStructBySn(sn)

	model := cpe.Model
	version := cpe.SoftwareVersion
	updatetime := cpe.EntryUpdateTime
	mastercpeip := cpe.MasterPopIP
	masterpopip := pv.GetPopStructById(cpe.MasterPopID).PopIP
	backupcpeip := cpe.BackupPopIP
	backuppopip := pv.GetPopStructById(cpe.BackupPopID).PopIP

	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyNexus(sn string) (string, string, string, string, string, string, string) {
	cpe := cn.GetCpeStructBySn(sn)

	model := cpe.Model
	version := cpe.SoftwareVersion
	updatetime := cpe.EntryUpdateTime
	mastercpeip := cpe.MasterEntryIP
	masterpopip := pn.GetPopStructById(cpe.MasterEntryID).EntryIP
	backupcpeip := cpe.BackupEntryIP
	backuppopip := pn.GetPopStructById(cpe.BackupEntryID).EntryIP

	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyWatsons(sn string) (string, string, string, string, string, string, string) {
	cpe := cw.GetCpeStructBySn(sn)

	model := cpe.Model
	version := cpe.SoftwareVersion
	updatetime := cpe.EntryUpdateTime
	mastercpeip := cpe.MasterEntryIP
	masterpopip := pw.GetPopStructById(cpe.MasterEntryID).EntryIP
	backupcpeip := cpe.BackupEntryIP
	backuppopip := pw.GetPopStructById(cpe.BackupEntryID).EntryIP

	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyWatsonsHa(sn string) (string, string, string, string, string, string, string) {
	cpe := ch.GetCpeStructBySn(sn)

	model := cpe.Model
	version := cpe.SoftwareVersion
	updatetime := cpe.EntryUpdateTime
	mastercpeip := cpe.MasterEntryIP
	masterpopip := ph.GetPopStructById(cpe.MasterEntryID).EntryIP
	backupcpeip := cpe.BackupEntryIP
	backuppopip := ph.GetPopStructById(cpe.BackupEntryID).EntryIP

	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}

func getCpebyZeratul(sn string) (string, string, string, string, string, string, string) {
	spe := cz.GetCpeStructBySn(sn)
	model := spe.Model
	version := spe.SoftwareVersion
	updatetime  := spe.PopUpdateTime
	mastercpeip := spe.MasterPopIP
	masterpopip := pz.GetPopStructById(spe.MasterPopID).EntryIP
	backupcpeip := spe.BackupPopIP
	backuppopip := pz.GetPopStructById(spe.BackupPopID).EntryIP

	return model, version, updatetime, masterpopip, mastercpeip, backuppopip, backupcpeip
}