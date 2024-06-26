package net

import (
	"fmt"
	"time"
	"net/http"
	"io"
	"encoding/json"
	"aixc/model/pop"
	. "aixc/tools"

)
var (
	pv pop.Valor
	py pop.Valor
	pw pop.Watsons
	ph pop.WatsonsHa
	pz pop.Zeratul
)

func getPopBytes(TOKEN, URL string) ([]byte, error) {
	Unix := TimeUnix(time.Now())
	popURL := fmt.Sprintf("%s?access_token=%s&_=%d", URL, TOKEN, Unix)

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

func getPopValorData(TOKEN, URL string) error {
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

func getPopYifengData(TOKEN, URL string) error {
	bytes, err := getPopBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &py); err != nil {
		return err
	}
	return nil
}

func getPopWatsonsData(TOKEN, URL string) error {
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

func getPopWatsonsHaData(TOKEN, URL string) error {
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

func getPopZeratulData(TOKEN, URL string) error {
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
