package net

import (
	"aixc/model/dve"
	tool "aixc/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	dv dve.Valor
	dn dve.Nexus
	dw dve.Watsons
	dh dve.WatsonsHa
	dz dve.Zeratul
)

func getDveBytes(TOKEN, URL string) ([]byte, error) {
	Unix := tool.TimeUnix(time.Now())
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

/*
http://internal.oss.7x-networks.net/matrix
/valor/device/cpes
/540
/configuration/enable_remote?
flag=true&
access_token=706ef88c-fb97-4cb0-984d-dae945386c8b
*/

func OpenValorRemote(SN, URL, TOKEN string) bool {
	dve := dv.GetDveStructBySn(SN)
	if dve.IsOnline() {
		if dve.EnableRemote {
			return true
		}
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "true", TOKEN)
		fmt.Println(OpenRemoteURL)
		// res, err := http.MethodPut
		return true
	}
	return false
}

func CloseValorRemote(SN, TOKEN, URL string) bool {
	dve := dv.GetDveStructBySn(SN)
	if dve.IsOnline() {
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "false", TOKEN)
		fmt.Println(OpenRemoteURL)
		return true
	}
	return false
}

func OpenZeratulRemote(SN, URL, TOKEN string) bool {
	dve := dz.GetDveStructBySn(SN)
	if dve.IsOnline() {
		if dve.EnableRemote {
			return true
		}
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "true", TOKEN)
		fmt.Println(OpenRemoteURL)
		// res, err := http.MethodPut
		return true
	}
	return false
}

func CloseZeratulRemote(SN, TOKEN, URL string) bool {
	dve := dz.GetDveStructBySn(SN)
	if dve.IsOnline() {
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "false", TOKEN)
		fmt.Println(OpenRemoteURL)
		return true
	}
	return false

}

func OpenWatonsRemote(SN, URL, TOKEN string) bool {
	dve := dw.GetDveStructBySn(SN)
	if dve.IsOnline() {
		if dve.SupportRemote {
			return true
		}
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "true", TOKEN)
		fmt.Println(OpenRemoteURL)
		// res, err := http.MethodPut
		return true
	}
	return false
}

func CloseWatonsRemote(SN, TOKEN, URL string) bool {
	dve := dw.GetDveStructBySn(SN)
	if dve.IsOnline() {
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "false", TOKEN)
		fmt.Println(OpenRemoteURL)
		return true
	}
	return false

}

func OpenWatonsHaRemote(SN, URL, TOKEN string) bool {
	dve := dh.GetDveStructBySn(SN)
	if dve.IsOnline() {
		if dve.SupportRemote {
			return true
		}
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "true", TOKEN)
		fmt.Println(OpenRemoteURL)
		// res, err := http.MethodPut
		return true
	}
	return false
}

func CloseWatonsHaRemote(SN, TOKEN, URL string) bool {
	dve := dh.GetDveStructBySn(SN)
	if dve.IsOnline() {
		OpenRemoteURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(URL, dve.ID), "false", TOKEN)
		fmt.Println(OpenRemoteURL)
		return true
	}
	return false
}

// 获取相同企业号的的UCPE
func getSnsByModeEn(mode, enterprise string) []string {
	var sns []string
	switch mode {
	case "valor":
		{
			sns = dv.GetCpesByEnterprise(enterprise)
		}
	case "tassadar":
		{
			sns = dz.GetCpesByEnterprise(enterprise)
		}
	case "nexus":
		{
			sns = dn.GetCpesByEnterprise(enterprise)
		}
	}
	return sns
}
