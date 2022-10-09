package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"aixc/model/cpe"
	"aixc/model/opt"
)

var (
	op opt.Operation

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

	bytes, err := ioutil.ReadAll(res.Body)
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

func getValorData(TOKEN, URL string) error {
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

func getNexusData(TOKEN, URL string) error {
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

func getWatsonsData(TOKEN, URL string) error {
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

func getWatsonsHaData(TOKEN, URL string) error {
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

func getZeratulData(TOKEN, URL string) error {
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
