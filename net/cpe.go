package net

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"aixc/model"
	"aixc/model/cpe"
)

var (
	op model.Operation

	cv cpe.Valor     //valor
	cn cpe.Nexus     //nexus
	cw cpe.Watsons   //watsons
	ch cpe.WatsonsHa //watsonsha
	cz cpe.Zeratul   //zeratul
)

func getCpeBytes(TOKEN, URL string) ([]byte, error) {
	Unix := timeUnix(time.Now())
	cpeURL := fmt.Sprintf("%saccess_token=%s&_=%d", URL, TOKEN, Unix)

	res, err := http.Get(cpeURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getOperationData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &op); err != nil {
		return err
	}
	return nil
}

func getCpeValorData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &cv); err != nil {
		return err
	}
	return nil
}

func getCpeNexusData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &cn); err != nil {
		return err
	}
	return nil
}

func getCpeWatsonsData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &cw); err != nil {
		return err
	}
	return nil
}

func getCpeWatsonsHaData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &ch); err != nil {
		return err
	}
	return nil
}

func getCpeZeratulData(TOKEN, URL string) error {
	body, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(body, &cz); err != nil {
		return err
	}

	return nil
}

func getCpebyValor(sn string) []string {
	cpe := cv.GetCpeStructBySn(sn)
	return []string {
		Cyan(sn),
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
		Cyan(sn),
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
		Cyan(sn),
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
		Cyan(sn),
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
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetPopStructById(spe.MasterPopID).EntryIP,
		spe.MasterPopIP,
		pz.GetPopStructById(spe.BackupPopID).EntryIP,
		spe.BackupPopIP,
	}
}

func cpeMaxVersionValor() string { return cv.MaxVersion()}
func cpeMaxVersionNexus() string { return cn.MaxVersion()}
func cpeMaxVersionWatsons() string { return cw.MaxVersion()}
func cpeMaxVersionWatsonsHa() string { return ch.MaxVersion()}
func cpeMaxVersionZeratul() string { return cz.MaxVersion()}


func ucpeInfoValor(sn string) []string {
	cpe := cv.GetCpeStructBySn(sn)
	dve := dv.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pv.GetPopStructById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		pv.GetPopStructById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		cpe.Alias,
	}
}

func ucpeInfoNexus(sn string) []string {
	cpe := cn.GetCpeStructBySn(sn)
	dve := dn.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		cpe.Model,cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pn.GetPopStructById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		pn.GetPopStructById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		cpe.Alias,
	}
}

func ucpeInfoWatsons(sn string) []string {
	cpe := cw.GetCpeStructBySn(sn)
	dve := dw.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pw.GetPopStructById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		pw.GetPopStructById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
		strconv.Itoa(dve.ServerPort),
		"watsons",
		cpe.Alias,

	}
}

func ucpeInfoWatsonsHa(sn string) []string {
	cpe := ch.GetCpeStructBySn(sn)
	dve := dh.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		ph.GetPopStructById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		ph.GetPopStructById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
		strconv.Itoa(dve.ServerPort),
		"watsonsha",
		cpe.Alias,
	}
}

func ucpeInfoZeratul(sn string) []string {
	spe := cz.GetCpeStructBySn(sn)
	dve := dz.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetPopStructById(spe.MasterPopID).EntryIP,
		spe.MasterPopIP,
		pz.GetPopStructById(spe.BackupPopID).EntryIP,
		spe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		spe.Alias,
	}
}

// show 企业号
func EucpeInfoValor(sn string) []string {
	cpe := cv.GetCpeStructBySn(sn)
	dve := dv.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pv.GetPopStructById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		pv.GetPopStructById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		cpe.Alias,
	}
}

func EucpeInfoNexus(sn string) []string {
	cpe := cn.GetCpeStructBySn(sn)
	dve := dn.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		cpe.Model,cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pn.GetPopStructById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		pn.GetPopStructById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
		strconv.Itoa(dve.ServerPort),
		cpe.Alias,
	}
}

func EucpeInfoZeratul(sn string) []string {
	spe := cz.GetCpeStructBySn(sn)
	dve := dz.GetDveStructBySn(sn)
	return []string {
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetPopStructById(spe.MasterPopID).EntryIP,
		spe.MasterPopIP,
		pz.GetPopStructById(spe.BackupPopID).EntryIP,
		spe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		spe.Alias,
	}
}

func allValor(sn string) [][]string {

	cpes := make([][]string, 0)

	if sn == "ALL" {
		for _, sn := range cz.SNs() {
			cpe := cv.GetCpeStructBySn(sn)
			dve := dv.GetDveStructBySn(sn)
			ucpe = []string {
				Cyan(sn),
				cpe.Model,
				cpe.SoftwareVersion,
				cpe.EntryUpdateTime,
				pv.GetPopStructById(cpe.MasterPopID).PopIP,
				cpe.MasterPopIP,
				pv.GetPopStructById(cpe.BackupPopID).PopIP,
				cpe.BackupPopIP,
				strconv.Itoa(dve.ServerPort),
				dve.Customer.Name,
				cpe.Alias,
			}
			// ucpe.Null().Version(MaxVersion).Time()
			cpes = append(cpes, ucpe)
		}
	}
	return cpes
}

func allZeratul(sn string) [][]string {
	cpes := make([][]string, 0)
	if sn == "ALL" {
		for _, sn := range cz.SNs() {
			spe := cz.GetCpeStructBySn(sn)
			dve := dz.GetDveStructBySn(sn)
			ucpe = []string {
				Cyan(sn),
				spe.Model,
				spe.SoftwareVersion,
				spe.PopUpdateTime,
				pz.GetPopStructById(spe.MasterPopID).EntryIP,
				spe.MasterPopIP,
				pz.GetPopStructById(spe.BackupPopID).EntryIP,
				spe.BackupPopIP,
				strconv.Itoa(dve.ServerPort),
				dve.Customer.Name,
				spe.Alias,
			}
			// ucpe.Null().Version(MaxVersion).Time()
			cpes = append(cpes, ucpe)
		}
	}

	return cpes
}

func allWatsonsHa(sn string) [][]string {
	cpes := make([][]string, 0)
	if sn == "ALL" {
		for _, sn := range ch.SNs() {
			spe := ch.GetCpeStructBySn(sn)
			dve := dh.GetDveStructBySn(sn)
			ucpe = []string {
				Cyan(sn),
				spe.Model,
				spe.SoftwareVersion,
				spe.EntryUpdateTime,
				ph.GetPopStructById(spe.MasterEntryID).EntryIP,
				spe.MasterEntryIP,
				ph.GetPopStructById(spe.BackupEntryID).EntryIP,
				spe.BackupEntryIP,
				strconv.Itoa(dve.ServerPort),
				"watsonsha",
				spe.Alias,
			}
			// ucpe.Null().Version(MaxVersion).Time()
			cpes = append(cpes, ucpe)
		}
	}

	return cpes
}