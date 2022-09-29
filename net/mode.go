package net


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
			mode = "watsonsha"
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

func getCpebyValor(sn string) []string {
	cpe := cv.GetCpeStructBySn(sn)
	return []string {
		cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pv.GetPopStructById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		pv.GetPopStructById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
	}
}

func getCpebyNexus(sn string) []string {
	cpe := cn.GetCpeStructBySn(sn)

	return []string {
		cyan(sn),
		cpe.Model,cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pn.GetPopStructById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		pn.GetPopStructById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
	}
}

func getCpebyWatsons(sn string) []string {
	cpe := cw.GetCpeStructBySn(sn)

	return []string {
		cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pw.GetPopStructById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		pw.GetPopStructById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
	}
}

func getCpebyWatsonsHa(sn string) []string {
	cpe := ch.GetCpeStructBySn(sn)

	return []string {
		cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		ph.GetPopStructById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		ph.GetPopStructById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
	}
}

func getCpebyZeratul(sn string) []string {
	spe := cz.GetCpeStructBySn(sn)
	return []string {
		cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetPopStructById(spe.MasterPopID).EntryIP,
		spe.MasterPopIP,
		pz.GetPopStructById(spe.BackupPopID).EntryIP,
		spe.BackupPopIP,
	}
}