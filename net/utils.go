package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func timeUnix(e time.Time) int64 {
	return e.UnixNano() / 1e6
}

func getResponseToken(URL string) (string, error) {

	// tokenurl := "http://internal.oss.7x-networks.net/matrix/oauth/token"

	data := make(url.Values)
	data["username"] = []string{"matrix"}
	data["password"] = []string{"c8d064e2ad4670f418ba02ef342b33d1"}
	data["client_id"] = []string{"browser"}
	data["client_secret"] = []string{"b7n3i7kzg22y3p035rw3rd9sfzvs4cv0"}
	data["grant_type"] = []string{"password"}
	
	res, err := http.PostForm(URL, data)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	token := result["access_token"].(string)
	
	return token, nil
}


func useMapArrgetMode(marr []map[string]string) string {
	arr := make([]string, 0)
	var m string
	for _, mp := range marr {
		for k, v := range mp {
			if v == "Yes" && k != "" {
				arr = append(arr, k)
				break
			}
		}
	}
	
	if len(arr) >= 2 {
		for _, i := range arr {
			m = "unknown"
			if i == "valor" {
				m = "valor"
				break
			}
		}
	} else if len(arr) == 1 {
		m = arr[0]
	} else {
		m = "unknown"
	}
	return m
}