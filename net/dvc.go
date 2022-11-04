package net

import (
	"fmt"
	"time"
	"io"
	"net/http"
	"encoding/json"

	"aixc/model/dvc"
)

var (
	dv dvc.Valor
	dn dvc.Nexus
	dw dvc.Watsons
	dh dvc.WatsonsHa
	dz dvc.Zeratul
)

func getDvcBytes(TOKEN, URL string) ([]byte, error) {
	Unix := timeUnix(time.Now())
	dvcURL := fmt.Sprintf("%saccess_token=%s&_=%d", URL, TOKEN, Unix)

	res, err := http.Get(dvcURL)
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


func getDvcNexusData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dn); err != nil {
		return err
	}
	return nil
}

func getDvcValorData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dv); err != nil {
		return err
	}
	return nil
}

func getDvcZeratulData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dz); err != nil {
		return err
	}
	return nil
}

func getDvcWatsonsData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dw); err != nil {
		return err
	}
	return nil
}

func getDvcWatsonsHaData(TOKEN, URL string) error {
	bytes, err := getCpeBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dh); err != nil {
		return err
	}
	return nil
}
