package net

import (
	"os"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/yaml.v3"

	"aixc/conf"
	"aixc/struct/opt"
	"aixc/struct/cpe"
	"aixc/struct/dvc"
	"aixc/struct/pop"
)

var (
	rules conf.Rules
	token string
	err   error

	opo opt.Operation

	cv  cpe.Valor     //valor
	cn  cpe.Nexus     //nexus
	cw  cpe.Watsons   //watsons
	ch  cpe.WatsonsHa //watsonsha
	cz 	cpe.Zeratul   //zeratul

	pv 	pop.Valor
	pn	pop.Nexus
	pw	pop.Watsons
	ph 	pop.WatsonsHa
	pz	pop.Zeratul

	dv  dvc.Valor
	dn  dvc.Nexus
	dw  dvc.Watsons
	dh  dvc.WatsonsHa
	dz  dvc.Zeratul
)

// 加载url路由规则
func init() {
	path := "/etc/aixc/url.rules"
	if err = loadURL(path); err != nil {
		os.Exit(9)
	}
	if err = getToken(rules.TokenRouteByMode()); err != nil {
		os.Exit(10)
	}
}

func loadURL(path string)  error {
	io, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(io, &rules) 
	if err != nil {
		return  err
	}
	return nil
}

// 获取token
func getToken(URL string) error {
	requestData := make(url.Values)
	requestData["username"]		 = []string{"matrix"}
	requestData["password"]		 = []string{"c8d064e2ad4670f418ba02ef342b33d1"}
	requestData["client_id"]	 = []string{"browser"}
	requestData["client_secret"] = []string{"b7n3i7kzg22y3p035rw3rd9sfzvs4cv0"}
	requestData["grant_type"]	 = []string{"password"}
	
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

	token = result["access_token"].(string)

	return nil
}

// 已知mode获取cpe,dvc,pop数据并放入到内存
func syncDataMemorybyMode(mode string) {
	cpeURL := rules.CpeRouteByMode(mode)
	popURL := rules.PopRouteByMode(mode)
	// dvcURL := rules.DeviceRouteByMode(mode)
	switch mode {
		case "valor":{
			cv, err = getValorData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pv, err = getValorPopData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dv, err = getDvcValorData(token, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
		case "nexus":{
			cn, err = getNexusData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pn, err = getNexusEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dn, err = getDvcNexusData(token, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
		case "watsons":{
			cw, err = getWatsonsData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pw, err = getWatsonsEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dw, err = getDvcWatsonsData(token, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
		case "watsons_ha":{
			ch, err = getWatsonsHaData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			ph, err = getWatsonsHaEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dh, err = getDvcWatsonsHaData(token, dvcURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
		case "tassadar":{
			cz, err = getZeratulData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pz, err = getZeratulPopData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dz, err = getDvcZeratulData(token, dvcURL)
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
		case "valor":{
			cv, err = getValorData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pv, err = getValorPopData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dv, err = getDvcValorData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			if ok, _ := cv.IsSn(sn); ok != false {
				return true
			}
		}
		case "nexus":{
			cn, err = getNexusData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pn, err = getNexusEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dn, err = getDvcNexusData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			if ok, _ := cn.IsSn(sn); ok != false {
				return true
			}
		}
		case "watsons":{
			cw, err = getWatsonsData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pw, err = getWatsonsEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dw, err = getDvcWatsonsData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			if ok, _ := cw.IsSn(sn); ok != false {
				return true
			}
		}
		case "watsons_ha":{
			ch, err = getWatsonsHaData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			ph, err = getWatsonsHaEntryData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dh, err = getDvcWatsonsHaData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			if ok, _ := ch.IsSn(sn); ok != false {
				return true
			}
		}
		case "tassadar":{
			cz, err = getZeratulData(token, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			pz, err = getZeratulPopData(token, popURL)
			if err != nil {
				os.Exit(13)
			}
			// dz, err = getDvcZeratulData(token, dvcURL)
			if err != nil {
				os.Exit(15)
			}
			if ok, _ := cz.IsSn(sn); ok != false {
				return true
			}
		}
	}
	return false
}

func getModebySevenSn(sn string) string {
	if opo, err = getOperationData(token, rules.OperationRouteByMode()); err != nil {
		os.Exit(11)
	}
	return opo.SnInMode(sn)
}