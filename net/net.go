package net

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"

	"aixc/conf"
	"aixc/struct/opt"
)

var (
	rules conf.Rules
	token string
	err   error

	opo opt.Operation
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
		fmt.Println(err.Error())
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