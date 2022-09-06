package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"aixc/struct/cpe"
	"aixc/struct/opt"
	"time"
)

func getCpeBytes(token, url string) ([]byte, error) {
	Unix := timeUnix(time.Now())
	cpeURL := fmt.Sprintf("%saccess_token=%s&_=%d", url, token, Unix)

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

func getOperationData(token, url string) (opt.Operation, error) {
	var operation opt.Operation
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return operation, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &operation); err != nil {
		return operation, err
	}
	return operation, nil
}

func getValorData(token, url string) (cpe.Valor, error) {
	var valor cpe.Valor
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &valor); err != nil {
		return nil, err
	}
	return valor, nil
}

func getNexusData(token, url string) (cpe.Nexus, error) {
	var nexus cpe.Nexus
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

func getWatsonsData(token, url string) (cpe.Watsons, error) {
	var ws cpe.Watsons
	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &ws); err != nil {
		return nil, err
	}
	return ws, nil
}

func getWatsonsHaData(token, url string) (cpe.WatsonsHa, error) {
	var whs cpe.WatsonsHa

	bytes, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &whs); err != nil {
		return nil, err
	}
	return whs, nil
}

func getZeratulData(token,url string)  (cpe.Zeratul, error) {
	var zeratul cpe.Zeratul
	
	body, err := getCpeBytes(token, url)
	if err != nil {
		return nil, err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(body, &zeratul); err != nil {
		return nil, err
	}

	return zeratul, nil
}