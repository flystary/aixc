package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	bytes, err := ioutil.ReadAll(res.Body)
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

func getValorData(TOKEN, URL string) error {
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

func getNexusData(TOKEN, URL string) error {
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

func getWatsonsData(TOKEN, URL string) error {
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

func getWatsonsHaData(TOKEN, URL string) error {
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

func getZeratulData(TOKEN, URL string) error {
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