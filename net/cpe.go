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
	. "aixc/tools"
	. "aixc/tools/color"
)

var (
	op model.Operation

	cv cpe.Valor     //valor
	cy cpe.Valor     //valor
	cw cpe.Watsons   //watsons
	ch cpe.WatsonsHa //watsonsha
	cz cpe.Zeratul   //zeratul
)

func getCpeBytes(TOKEN, URL string) ([]byte, error) {
	Unix := TimeUnix(time.Now())
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

func getCpeYifengData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &cy); err != nil {
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
	return []string{
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

func getCpebyYifeng(sn string) []string {
	cpe := cy.GetCpeStructBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		py.GetPopStructById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		py.GetPopStructById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
	}
}

func getCpebyWatsons(sn string) []string {
	cpe := cw.GetCpeStructBySn(sn)
	return []string{
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
	return []string{
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
	return []string{
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetPopStructById(spe.MasterPopID).PopIP,
		spe.MasterPopIP,
		pz.GetPopStructById(spe.BackupPopID).PopIP,
		spe.BackupPopIP,
	}
}

func cpeMaxVersionValor() string     { return cv.MaxVersion() }
func cpeMaxVersionYifeng() string    { return cy.MaxVersion() }
func cpeMaxVersionWatsons() string   { return cw.MaxVersion() }
func cpeMaxVersionWatsonsHa() string { return ch.MaxVersion() }
func cpeMaxVersionZeratul() string   { return cz.MaxVersion() }

func uCPEInfoValor(sn string) []string {
	cpe := cv.GetCpeStructBySn(sn)
	dve := dv.GetDveStructBySn(sn)
	return []string{
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

func uCPEInfoWatsons(sn string) []string {
	cpe := cw.GetCpeStructBySn(sn)
	dve := dw.GetDveStructBySn(sn)
	return []string{
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
func uCPEInfoYifeng(sn string) []string {
	cpe := cy.GetCpeStructBySn(sn)
	dve := dy.GetDveStructBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		py.GetPopStructById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		py.GetPopStructById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		cpe.Alias,
	}
}

func uCPEInfoWatsonsHa(sn string) []string {
	cpe := ch.GetCpeStructBySn(sn)
	dve := dh.GetDveStructBySn(sn)
	return []string{
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

func uCPEInfoZeratul(sn string) []string {
	spe := cz.GetCpeStructBySn(sn)
	dve := dz.GetDveStructBySn(sn)
	return []string{
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetPopStructById(spe.MasterPopID).PopIP,
		spe.MasterPopIP,
		pz.GetPopStructById(spe.BackupPopID).PopIP,
		spe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		spe.Alias,
	}
}

// show 企业号
func EuCPEInfoValor(sn string) []string {
	cpe := cv.GetCpeStructBySn(sn)
	dve := dv.GetDveStructBySn(sn)
	return []string{
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

func EuCPEInfoYifeng(sn string) []string {
	cpe := cy.GetCpeStructBySn(sn)
	dve := dy.GetDveStructBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		py.GetPopStructById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		py.GetPopStructById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		cpe.Alias,
	}
}

func EuCPEInfoZeratul(sn string) []string {
	spe := cz.GetCpeStructBySn(sn)
	dve := dz.GetDveStructBySn(sn)
	return []string{
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetPopStructById(spe.MasterPopID).PopIP,
		spe.MasterPopIP,
		pz.GetPopStructById(spe.BackupPopID).PopIP,
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
			ucpe = []string{
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

func allYifeng(sn string) [][]string {

	cpes := make([][]string, 0)

	if sn == "ALL" {
		for _, sn := range cy.SNs() {
			cpe := cy.GetCpeStructBySn(sn)
			dve := dy.GetDveStructBySn(sn)
			ucpe = []string{
				Cyan(sn),
				cpe.Model,
				cpe.SoftwareVersion,
				cpe.EntryUpdateTime,
				py.GetPopStructById(cpe.MasterPopID).PopIP,
				cpe.MasterPopIP,
				py.GetPopStructById(cpe.BackupPopID).PopIP,
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
			ucpe = []string{
				Cyan(sn),
				spe.Model,
				spe.SoftwareVersion,
				spe.PopUpdateTime,
				pz.GetPopStructById(spe.MasterPopID).PopIP,
				spe.MasterPopIP,
				pz.GetPopStructById(spe.BackupPopID).PopIP,
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
			ucpe = []string{
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

// 获取相同model的的uCPE
func getSnsByModel(mode, model string) []string {
	var sns []string
	switch mode {
	case "valor":
		{
			sns = cv.GetCpesByModel(model)
		}
	case "yifeng":
		{
			sns = cy.GetCpesByModel(model)
		}
	case "tassadar":
		{
			sns = cz.GetCpesByModel(model)
		}
	case "watsons":
		{
			sns = cw.GetCpesByModel(model)
		}
	case "watsonsha":
		{
			sns = ch.GetCpesByModel(model)
		}
	}
	return sns
}

// 获取相同version的的uCPE
func getSnsByVersion(mode, version string) []string {
	var sns []string
	switch mode {
	case "valor":
		{
			sns = cv.GetCpesByVersion(version)
		}
	case "yifeng":
		{
			sns = cy.GetCpesByVersion(version)
		}
	case "tassadar":
		{
			sns = cz.GetCpesByVersion(version)
		}
	case "watsons":
		{
			sns = cw.GetCpesByVersion(version)
		}
	case "watsonsha":
		{
			sns = ch.GetCpesByVersion(version)
		}
	}
	return sns
}

// 获取相同Addr的的uCPE
func getSnsByPopAddr(mode, addr string) []string {
	var id = 0
	var sns []string
	switch mode {
	case "valor":
		{
			id = pv.GetIdByAddr(addr)
			sns = cv.GetCpesByPopId(id)
		}
	case "yifeng":
		{
			id = py.GetIdByAddr(addr)
			sns = cy.GetCpesByPopId(id)
		}
	case "tassadar":
		{
			id = pz.GetIdByAddr(addr)
			sns = cz.GetCpesByPopId(id)
		}
	case "watsons":
		{
			id = pw.GetIdByAddr(addr)
			sns = cw.GetCpesByPopId(id)
		}
	case "watsonsha":
		{
			id = ph.GetIdByAddr(addr)
			sns = ch.GetCpesByPopId(id)
		}
	}
	return sns
}
