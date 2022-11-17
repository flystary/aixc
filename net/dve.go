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
	dw dve.Watsons
	dh dve.WatsonsHa
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

func OpenValorEnableRemote(TOKEN, URL string, report int)  {
	OpenRemoteURL := fmt.Sprintf("%s/%d/configuration/enable_remote?flag=true&access_token=%s", URL, report, TOKEN)
	// http://internal.oss.7x-networks.net/matrix/valor/device/cpes/334/configuration/enable_remote?flag=true&access_token=79fda75a-70f1-4c82-8ad1-539e287b3dc7
	fmt.Println(OpenRemoteURL)
}

func CloseEnableRemote(TOKEN, URL string) {

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

func getDveWatsonsData(TOKEN, URL string) error {
	bytes, err := getDveBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dw); err != nil {
		return err
	}
	return nil
}

func getDveWatsonsHaData(TOKEN, URL string) error {
	bytes, err := getDveBytes(TOKEN, URL)
	if err != nil {
		return err
	}
	// Unmarshal json数据
	if err = json.Unmarshal(bytes, &dh); err != nil {
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