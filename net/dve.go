package net

import (
	"fmt"
	"time"
	"io"
	"net/http"
	"encoding/json"

	"aixc/model/dve"
)

var (
	dv dve.Valor
	dn dve.Nexus
	dz dve.Zeratul
)

func getDveBytes(TOKEN, URL string) ([]byte, error) {
	Unix := timeUnix(time.Now())
	dveURL := fmt.Sprintf("%saccess_token=%s&_=%d", URL, TOKEN, Unix)

	res, err := http.Get(dveURL)
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

func getDveNexusData(TOKEN, URL string) error {
	bytes, err := getDveBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dn); err != nil {
		return err
	}
	return nil
}

func getDveValorData(TOKEN, URL string) error {
	bytes, err := getDveBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dv); err != nil {
		return err
	}
	return nil
}

func getDveZeratulData(TOKEN, URL string) error {
	bytes, err := getDveBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dz); err != nil {
		return err
	}
	return nil
}

func getSnsByMode(mode,enterprise string) []string {
	var sns []string
	switch mode {
		case "valor":{
			sns = dv.GetCpesByEnterprise(enterprise)
		}
		case "tassadar":{
			sns = dz.GetCpesByEnterprise(enterprise)
		}
		case "nexus":{
			sns = dn.GetCpesByEnterprise(enterprise)
		}
	}
	return sns
}