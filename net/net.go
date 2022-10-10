package net

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"

	"aixc/conf"
)

var (
	rules conf.Rules
	// TOKEN 密钥
	TOKEN string
	err   error

)

// 加载url路由规则
func init() {
	path := "/etc/xc/url.rules"
	if err = loadURL(path); err != nil {
		os.Exit(9)
	}
	if err = getToken(rules.TokenRouteByMode()); err != nil {
		os.Exit(10)
	}
}

func loadURL(path string) error {
	io, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(io, &rules)
	if err != nil {
		return err
	}
	return nil
}

// 获取token
func getToken(URL string) error {
	requestData := make(url.Values)
	requestData["username"] = []string{"matrix"}
	requestData["password"] = []string{newMD5(newMD5("4A9sOpYL"))}
	requestData["client_id"] = []string{"browser"}
	requestData["client_secret"] = []string{"b7n3i7kzg22y3p035rw3rd9sfzvs4cv0"}
	requestData["grant_type"] = []string{"password"}

	res, err := http.PostForm(URL, requestData)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	TOKEN = result["access_token"].(string)
	return nil
}

// 已知mode获取cpe,dvc,pop数据并放入到内存
func syncDataMemorybyMode(mode string) {
	cpeURL := rules.CpeRouteByMode(mode)
	popURL := rules.PopRouteByMode(mode)
	// dvcURL := rules.DeviceRouteByMode(mode)
	switch mode {
	case "valor":
		{
			err = getValorData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getValorPopData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcValorData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	case "nexus":
		{
			err = getNexusData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getNexusEntryData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcNexusData(token, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	case "watsons":
		{
			err = getWatsonsData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getWatsonsEntryData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcWatsonsData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	case "watsonsha":
		{
			err = getWatsonsHaData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getWatsonsHaEntryData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcWatsonsHaData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	case "tassadar":
		{
			err = getZeratulData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getZeratulPopData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcZeratulData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	}
}

// 根据sn和mode获取cpe,dvc,pop数据并放入到内存
func syncDataMemorybySnMode(sn, mode string) bool {
	cpeURL := rules.CpeRouteByMode(mode)
	popURL := rules.PopRouteByMode(mode)
	// dvcURL := rules.DeviceRouteByMode(mode)
	switch mode {
	case "valor":
		{
			err = getValorData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getValorPopData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcValorData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if ok, _ := cv.IsSn(sn); ok != false {
				return true
			}
		}
	case "nexus":
		{
			err = getNexusData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getNexusEntryData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcNexusData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if ok, _ := cn.IsSn(sn); ok != false {
				return true
			}
		}
	case "watsons":
		{
			err = getWatsonsData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getWatsonsEntryData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcWatsonsData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if ok, _ := cw.IsSn(sn); ok != false {
				return true
			}
		}
	case "watsonsha":
		{
			err = getWatsonsHaData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getWatsonsHaEntryData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcWatsonsHaData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if ok, _ := ch.IsSn(sn); ok != false {
				return true
			}
		}
	case "tassadar":
		{
			err = getZeratulData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getZeratulPopData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDvcZeratulData(TOKEN, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if ok, _ := cz.IsSn(sn); ok != false {
				return true
			}
		}
	}
	return false
}

func getModebySevenSn(sn string) string {
	if err = getOperationData(TOKEN, rules.OperationRouteByMode()); err != nil {
		os.Exit(11)
	}
	return op.SnInMode(sn)
}
