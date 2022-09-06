package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"aixc/struct/dvc"
	"time"
)

func getDvcBytes(token, url string) ([]byte, error) {
	Unix := timeUnix(time.Now())
	dvcURL := fmt.Sprintf("%saccess_token=%s&_=%d", url, token, Unix)

	res, err := http.Get(dvcURL)
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

func getDvcNexusData(token, url string) (dvc.Nexus, error) {
	var nexus dvc.Nexus
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &nexus); err != nil {
		return nexus, err
	}
	return nexus, nil
}

func getDvcValorData(token, url string) (dvc.Valor, error) {
	var valor dvc.Valor
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &valor); err != nil {
		return valor, err
	}
	return valor, nil
}

func getDvcZeratulData(token, url string) (dvc.Zeratul, error) {
	var zeratul dvc.Zeratul
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &zeratul); err != nil {
		return zeratul, err
	}
	return zeratul, nil
}

func getDvcWatsonsData(token, url string) (dvc.Watsons, error) {
	var watsons dvc.Watsons
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &watsons); err != nil {
		return watsons, err
	}
	return watsons, nil
}

func getDvcWatsonsHaData(token, url string) (dvc.WatsonsHa, error) {
	var watsonsha dvc.WatsonsHa
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &watsonsha); err != nil {
		return watsonsha, err
	}
	return watsonsha, nil
}
