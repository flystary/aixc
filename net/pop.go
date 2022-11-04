package net

import (
	"fmt"
	"time"
	"net/http"
	"io"
	"encoding/json"
	"aixc/model/pop"
)
var (
	pv pop.Valor
	pn pop.Nexus
	pw pop.Watsons
	ph pop.WatsonsHa
	pz pop.Zeratul
)

func getPopBytes(TOKEN, URL string) ([]byte, error) {
	Unix := timeUnix(time.Now())
	popURL := fmt.Sprintf("%saccess_token=%s&_=%d", URL, TOKEN, Unix)

	res, err := http.Get(popURL)
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

func getNexusEntryData(TOKEN, URL string) error {
	bytes, err := getPopBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &pn); err != nil {
		return err
	}
	return nil
}

func getValorPopData(TOKEN, URL string) error {
	bytes, err := getPopBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &pv); err != nil {
		return err
	}
	return nil
}

func getWatsonsEntryData(TOKEN, URL string) error {
	bytes, err := getPopBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &pw); err != nil {
		return err
	}
	return nil
}

func getWatsonsHaEntryData(TOKEN, URL string) error {
	bytes, err := getPopBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &ph); err != nil {
		return err
	}
	return nil
}

func getZeratulPopData(TOKEN, URL string) error {
	bytes, err := getPopBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &pz); err != nil {
		return err
	}
	return nil
}
