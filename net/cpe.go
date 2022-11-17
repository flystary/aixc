package net

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"aixc/model/cpe"
	"aixc/model"
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

func GetCpebyValor(sn string) []string {
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

func GetCpebyNexus(sn string) []string {
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

func GetCpebyWatsons(sn string) []string {
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

func GetCpebyWatsonsHa(sn string) []string {
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

func GetCpebyZeratul(sn string) []string {
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