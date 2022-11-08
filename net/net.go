package net

import (

	"fmt"
	"os"

	r "aixc/route"
)

var (
 	// TOKEN
	TOKEN  string
	err    error
	route  r.Route
	neter  Neter = &route
)

// Neter
type Neter interface {
	Router
	// Requester
	// Serializer
}

// Router
type Router interface {
	GetCpeFromRoute(mode string) string
	GetPopFromRoute(mode string) string
	GetDveFromRoute(mode string) string
	GetOperationFromRoute() string
	GetTokenFromRoute() string
}

// Requester
type Requester interface{
	GetToken(URL string) string
	// GetMode(service  *Service) *Service
	// GetBytes(service *Service) *Service
}

// Serializer 序列化
type Serializer interface{}

// 加载url路由规则
func init() {
	fileName := "route.yaml"
	path := fmt.Sprintf("/etc/aixc/%s", fileName)
	route = r.LoadRoute(path)
	// 获取token
	TOKEN = GetToken(neter.GetTokenFromRoute())
}

// 已知mode获取cpe,dve,pop数据并放入到内存
func SyncEnDataMemorybyMode(mode string) {
	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	dveURL := neter.GetDveFromRoute(mode)
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
			err = getDveValorData(TOKEN, dveURL)
			if err != nil {
				os.Exit(15)
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
			err = getDveZeratulData(TOKEN, dveURL)
			if err != nil {
				os.Exit(15)
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
			err = getDveNexusData(TOKEN, dveURL)
			if err != nil {
				os.Exit(15)
			}
		}
	}

}

// 已知mode获取cpe,dve,pop数据并放入到内存
func SyncDataMemorybyMode(mode string) {
	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	dveURL := neter.GetDveFromRoute(mode)
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
			err = getDveValorData(TOKEN, dveURL)
			if err != nil {
				os.Exit(15)
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
			err = getDveZeratulData(TOKEN, dveURL)
			if err != nil {
				os.Exit(15)
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
			// err = getDveWatsonsData(TOKEN, dveURL)
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
			// err = getDveWatsonsHaData(TOKEN, dveURL)
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
			err = getDveNexusData(TOKEN, dveURL)
			if err != nil {
				os.Exit(15)
			}
		}
	}

}

// 根据sn和mode获取cpe,dve,pop数据并放入到内存
func SyncDataMemorybySnMode(sn, mode string) bool {
	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	// dveURL := rules.DeviceRouteByMode(mode)
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
			// err = getDveValorData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if is, _ := cv.IsSn(sn); is{
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
			// err = getDveNexusData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if is, _ := cn.IsSn(sn); is{
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
			// err = getDveWatsonsData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if is, _ := cw.IsSn(sn); is{
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
			// err = getDveWatsonsHaData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if is, _ := ch.IsSn(sn); is{
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
			// err = getDveZeratulData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
			if is, _ := cz.IsSn(sn); is{
				return true
			}
		}
	}
	return false
}

func GetModebySevenSn(sn string) string {
	if err = getOperationData(TOKEN, neter.GetOperationFromRoute()); err != nil {
		os.Exit(11)
	}
	return op.SnInMode(sn)
}
