package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"

	"gopkg.in/yaml.v3"

	"aixc/conf"
)

var (
	rules *conf.Rules
	// 文件单例
	once  = &sync.Once{}
	// token单例
	onces = &sync.Once{}
 	// TOKEN 密钥
	TOKEN  string
	err    error
)

// 加载url路由规则
func init() {
	path := "/etc/xc/url.rules"
	loadURL(path)
	getToken(rules.TokenRoute())
}

func loadURL(path string) {
	once.Do(func() {
		io, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("Open File Error: %v", err)
		}
		err = yaml.Unmarshal(io, &rules)
		if err != nil {
			fmt.Printf("Unmarshal File Error: %v", err)
		}
	})
}

// 获取token
func getToken(URL string)  {
	var result = make(map[string]interface{})
	requestData := make(url.Values)
	requestData["username"] = []string{"matrix"}
	requestData["password"] = []string{newMD5(newMD5("4A9sOpYL"))}
	requestData["client_id"] = []string{"browser"}
	requestData["client_secret"] = []string{"b7n3i7kzg22y3p035rw3rd9sfzvs4cv0"}
	requestData["grant_type"] = []string{"password"}

	onces.Do(func() {
		res, err := http.PostForm(URL, requestData)
		if err != nil {
			fmt.Printf("Login Error: %v", err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("ReadAll IO Error: %v", err)
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Printf("Unmarshal body Error: %v", err)
		}
		TOKEN = result["access_token"].(string)
	})
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
	if err = getOperationData(TOKEN, rules.OperationRoute()); err != nil {
		os.Exit(11)
	}
	return op.SnInMode(sn)
}
