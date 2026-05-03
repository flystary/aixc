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
	cpe, _ := cv.Data.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pv.Data.GetById(cpe.MasterPopID).GetIP(),
		cpe.MasterPopIP,
		pv.Data.GetById(cpe.BackupPopID).GetIP(),
		cpe.BackupPopIP,
	}
}

func getCpebyYifeng(sn string) []string {
	cpe, _ := cy.Data.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		py.Data.GetById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		py.Data.GetById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
	}
}

func getCpebyWatsons(sn string) []string {
	cpe, _ := cw.Data.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pw.Data.GetById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		pw.Data.GetById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
	}
}

func getCpebyWatsonsHa(sn string) []string {
	cpe, _ := ch.Data.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		ph.Data.GetById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		ph.Data.GetById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
	}
}

func getCpebyZeratul(sn string) []string {
	spe, _ := cz.Data.GetBySn(sn)
	return []string{
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.Data.GetById(spe.MasterPopID).PopIP,
		spe.MasterPopIP,
		pz.Data.GetById(spe.BackupPopID).PopIP,
		spe.BackupPopIP,
	}
}

func cpeMaxVersionValor() string     { return cv.Data.MaxVersion() }
func cpeMaxVersionYifeng() string    { return cy.Data.MaxVersion() }
func cpeMaxVersionWatsons() string   { return cw.Data.MaxVersion() }
func cpeMaxVersionWatsonsHa() string { return ch.Data.MaxVersion() }
func cpeMaxVersionZeratul() string   { return cz.Data.MaxVersion() }

func uCPEInfoValor(sn string) []string {
	cpe, _ := cv.GetBySn(sn)
	dve := dv.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pv.Data.GetById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		pv.Data.GetById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		cpe.Alias,
	}
}

func uCPEInfoWatsons(sn string) []string {
	cpe, _ := cw.GetBySn(sn)
	dve := dw.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pw.GetById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		pw.GetById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
		strconv.Itoa(dve.ServerPort),
		"watsons",
		cpe.Alias,
	}
}
func uCPEInfoYifeng(sn string) []string {
	cpe, _ := cy.GetBySn(sn)
	dve := dy.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		py.GetById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		py.GetById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		cpe.Alias,
	}
}

func uCPEInfoWatsonsHa(sn string) []string {
	cpe, _ := ch.GetBySn(sn)
	dve := dh.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		ph.GetById(cpe.MasterEntryID).EntryIP,
		cpe.MasterEntryIP,
		ph.GetById(cpe.BackupEntryID).EntryIP,
		cpe.BackupEntryIP,
		strconv.Itoa(dve.ServerPort),
		"watsonsha",
		cpe.Alias,
	}
}

func uCPEInfoZeratul(sn string) []string {
	spe, _ := cz.GetBySn(sn)
	dve := dz.GetBySn(sn)
	return []string{
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetById(spe.MasterPopID).PopIP,
		spe.MasterPopIP,
		pz.GetById(spe.BackupPopID).PopIP,
		spe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		dve.Customer.Name,
		spe.Alias,
	}
}

// show 企业号
func EuCPEInfoValor(sn string) []string {
	cpe, _ := cv.GetBySn(sn)
	dve := dv.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		pv.GetById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		pv.GetById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		cpe.Alias,
	}
}

func EuCPEInfoYifeng(sn string) []string {
	cpe, _ := cy.GetBySn(sn)
	dve := dy.GetBySn(sn)
	return []string{
		Cyan(sn),
		cpe.Model,
		cpe.SoftwareVersion,
		cpe.EntryUpdateTime,
		py.GetById(cpe.MasterPopID).PopIP,
		cpe.MasterPopIP,
		py.GetById(cpe.BackupPopID).PopIP,
		cpe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		cpe.Alias,
	}
}

func EuCPEInfoZeratul(sn string) []string {
	spe, _ := cz.GetBySn(sn)
	dve := dz.GetBySn(sn)
	return []string{
		Cyan(sn),
		spe.Model,
		spe.SoftwareVersion,
		spe.PopUpdateTime,
		pz.GetById(spe.MasterPopID).PopIP,
		spe.MasterPopIP,
		pz.GetById(spe.BackupPopID).PopIP,
		spe.BackupPopIP,
		strconv.Itoa(dve.ServerPort),
		spe.Alias,
	}
}

func allValor(sn string) [][]string {

	cpes := make([][]string, 0)

	if sn == "ALL" {
		for _, sn := range cz.SNs() {
			cpe, _ := cv.GetBySn(sn)
			dve := dv.GetBySn(sn)
			ucpe = []string{
				Cyan(sn),
				cpe.Model,
				cpe.SoftwareVersion,
				cpe.EntryUpdateTime,
				pv.GetById(cpe.MasterPopID).PopIP,
				cpe.MasterPopIP,
				pv.GetById(cpe.BackupPopID).PopIP,
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
		for _, sn := range cy.Data.SNs() {
			cpe, _ := cy.Data.GetBySn(sn)
			dve := dy.Data.GetBySn(sn)
			ucpe = []string{
				Cyan(sn),
				cpe.Model,
				cpe.SoftwareVersion,
				cpe.EntryUpdateTime,
				py.Data.GetById(cpe.MasterPopID).PopIP,
				cpe.MasterPopIP,
				py.Data.GetById(cpe.BackupPopID).PopIP,
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
			spe, _ := cz.GetBySn(sn)
			dve := dz.GetBySn(sn)
			ucpe = []string{
				Cyan(sn),
				spe.Model,
				spe.SoftwareVersion,
				spe.PopUpdateTime,
				pz.Data.GetById(spe.MasterPopID).PopIP,
				spe.MasterPopIP,
				pz.Data.GetById(spe.BackupPopID).PopIP,
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
		for _, sn := range ch.Data.SNs() {
			spe, _ := ch.Data.GetBySn(sn)
			dve := dh.Data.GetBySn(sn)
			ucpe = []string{
				Cyan(sn),
				spe.Model,
				spe.SoftwareVersion,
				spe.EntryUpdateTime,
				ph.Data.GetById(spe.MasterEntryID).EntryIP,
				spe.MasterEntryIP,
				ph.Data.GetById(spe.BackupEntryID).EntryIP,
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
			sns = cv.Data.GetByModel(model)
		}
	case "yifeng":
		{
			sns = cy.Data.GetByModel(model)
		}
	case "tassadar":
		{
			sns = cz.Data.GetByModel(model)
		}
	case "watsons":
		{
			sns = cw.Data.GetByModel(model)
		}
	case "watsonsha":
		{
			sns = ch.Data.GetByModel(model)
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
			sns = cv.GetByVersion(version)
		}
	case "yifeng":
		{
			sns = cy.GetByVersion(version)
		}
	case "tassadar":
		{
			sns = cz.GetByVersion(version)
		}
	case "watsons":
		{
			sns = cw.GetByVersion(version)
		}
	case "watsonsha":
		{
			sns = ch.GetByVersion(version)
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
			id = pv.Data.GetIDByAddr(addr)
			sns = cv.Data.GetByPopId(id)
		}
	case "yifeng":
		{
			id = py.Data.GetIDByAddr(addr)
			sns = cy.Data.GetByPopId(id)
		}
	case "tassadar":
		{
			id = pz.GetIDByAddr(addr)
			sns = cz.Data.GetByPopID(id)
		}
	case "watsons":
		{
			id = pw.GetIDByAddr(addr)
			sns = cw.GetByPopID(id)
		}
	case "watsonsha":
		{
			id = ph.GetIDByAddr(addr)
			sns = ch.GetByPopID(id)
		}
	}
	return sns
}
