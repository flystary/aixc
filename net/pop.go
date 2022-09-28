package net

import (
	"aixc/struct/pop"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getPopBytes(token, url string) ([]byte, error) {
	Unix := timeUnix(time.Now())
	popURL := fmt.Sprintf("%saccess_token=%s&_=%d", url, token, Unix)

	res, err := http.Get(popURL)
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

func getNexusEntryData(token, url string) (pop.Nexus, error) {
	var nexus pop.Nexus
	bytes, err := getPopBytes(token, url)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &nexus); err != nil {
		return nil, err
	}
	return nexus, nil
}

func getValorPopData(token, url string) (pop.Valor, error) {
	var valor pop.Valor
	bytes, err := getPopBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &valor); err != nil {
		return nil, err
	}
	return valor, nil
}

func getWatsonsEntryData(token, url string) (pop.Watsons, error) {
	var watsons pop.Watsons
	bytes, err := getPopBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &watsons); err != nil {
		return watsons, err
	}
	return watsons, nil
}

func getWatsonsHaEntryData(token, url string) (pop.WatsonsHa, error) {
	var watsonsha pop.WatsonsHa
	bytes, err := getPopBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &watsonsha); err != nil {
		return nil, err
	}
	return watsonsha, nil
}

func getZeratulPopData(token, url string) (pop.Zeratul, error) {
	var zeratul pop.Zeratul
	bytes, err := getPopBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err := json.Unmarshal(bytes, &zeratul); err != nil {
		return nil, err
	}
	return zeratul, nil
}
