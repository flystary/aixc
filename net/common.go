package net

import (
	"aixc/model/cpe"
	"aixc/model/dve"
	"aixc/model/pop"
	. "aixc/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// T ~ cpe.Valor, pop.Watsons dev.Watsons
func FetchAny[T any](TOKEN, URL string, target *T) error {
	Unix := TimeUnix(time.Now())
	fullURL := fmt.Sprintf("%s?access_token=%s&_=%d", URL, TOKEN, Unix)

	res, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, target)
}

var CpeRegistry = make(map[string]cpe.CpeProvider)
var DveRegistry = make(map[string]dve.DveProvider)
var PopRegistry = make(map[string]pop.PopProvider)

func InitAllCpes(token string, urls map[string]string) {
	for mode, url := range urls {
		var err error
		var provider cpe.CpeProvider

		switch mode {
		case "valor", "yifeng":
			var d cpe.Valor
			err = FetchAny(token, url, &d)
			provider = d
		case "watsons":
			var d cpe.Watsons
			err = FetchAny(token, url, &d)
			provider = d
		case "watsonsha":
			var d cpe.WatsonsHa
			err = FetchAny(token, url, &d)
			provider = d
		case "zeratul":
			var d cpe.Zeratul
			err = FetchAny(token, url, &d)
			provider = d
		}

		if err != nil {
			fmt.Printf("Failed to fetch %s: %v\n", mode, err)
			continue
		}

		// 统一注册
		CpeRegistry[mode] = provider
	}
}

func InitAllDves(token string, urls map[string]string) {
	for mode, url := range urls {
		var err error
		var provider dve.DveProvider

		switch mode {
		case "valor", "yifeng":
			var d dve.Valor
			err = FetchAny(token, url, &d)
			provider = d
		case "watsons":
			var d dve.Watsons
			err = FetchAny(token, url, &d)
			provider = d
		case "watsonsha":
			var d dve.WatsonsHa
			err = FetchAny(token, url, &d)
			provider = d
		case "zeratul":
			var d dve.Zeratul
			err = FetchAny(token, url, &d)
			provider = d
		}

		if err != nil {
			fmt.Printf("Failed to fetch %s: %v\n", mode, err)
			continue
		}

		// 统一注册
		DveRegistry[mode] = provider
	}
}

func InitAllPops(token string, urls map[string]string) {
	for mode, url := range urls {
		var err error
		var provider pop.PopProvider

		switch mode {
		case "valor", "yifeng":
			var d pop.Valor
			err = FetchAny(token, url, &d)
			provider = d
		case "watsons":
			var d pop.Watsons
			err = FetchAny(token, url, &d)
			provider = d
		case "watsonsha":
			var d pop.WatsonsHa
			err = FetchAny(token, url, &d)
			provider = d
		case "zeratul":
			var d pop.Zeratul
			err = FetchAny(token, url, &d)
			provider = d
		}

		if err != nil {
			fmt.Printf("Failed to fetch %s: %v\n", mode, err)
			continue
		}

		// 统一注册
		PopRegistry[mode] = provider
	}
}

// SetRemoteStatus 统一远程开关
func SetRemoteStatus(mode, sn, token, actionUrl, targetStatus string) bool {
	provider, ok := DveRegistry[mode]
	if !ok {
		fmt.Printf("未找到模式 %s 的注册信息\n", mode)
		return false
	}

	device := provider.GetCollection().GetBySn(sn)
	if device == nil {
		return false
	}

	if device.IsOnline() {
		currentStatus := "false"
		if device.IsRemoteEnabled() {
			currentStatus = "true"
		}
		if currentStatus == targetStatus {
			return true
		}
		// send action
		fullURL := fmt.Sprintf("%s%s&access_token=%s", fmt.Sprintf(actionUrl, device), targetStatus, token)

		fmt.Println("正在发送远程控制请求:", fullURL)
		return true
	}

	return false
}
