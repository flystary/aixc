package net

import (
	"fmt"
	"os"
	"sync"

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
func SyncDataMemorybyMode(mode string) {
	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	// dveURL := neter.GetDveFromRoute(mode)
	switch mode {
	case "valor":
		{
			err = getCpeValorData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopValorData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDveValorData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	case "tassadar":
		{
			err = getCpeZeratulData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopZeratulData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDveZeratulData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	case "watsons":
		{
			err = getCpeWatsonsData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopWatsonsData(TOKEN, popURL)
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
			err = getCpeWatsonsHaData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopWatsonsHaData(TOKEN, popURL)
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
			err = getCpeNexusData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopNexusData(TOKEN, popURL)
			if err != nil {
				os.Exit(13)
			}
			// err = getDveNexusData(TOKEN, dveURL)
			// if err != nil {
			// 	os.Exit(15)
			// }
		}
	}
}

// 已知sn和随机mode获取cpe,dve,pop数据并放入到内存
func SyncDataMemorybySnMode(sn, mode string) bool {
	wg  := &sync.WaitGroup{}
	num := 2
	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	// dveURL := neter.GetDveFromRoute(mode)
	switch mode {
	case "valor":
		{
			var  err1  error
			wg.Add(num)
			go func ()  {
				err1 = getCpeValorData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func ()  {
				err1 = getPopValorData(TOKEN, popURL)
				wg.Done()
			}()
			// go func ()  {
			// 	err1 =getDveValorData(TOKEN, dveURL)
			// 	wg.Done()
			// }()
			wg.Wait()
			if err1 != nil {
				os.Exit(12)
			}
			if is, _ := cv.IsSn(sn); is {
				return true
			}
		}
	case "nexus":
		{
			var  err1  error
			wg.Add(num)
			go func ()  {
				err1 = getCpeNexusData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func ()  {
				err1 = getPopNexusData(TOKEN, popURL)
				wg.Done()
			}()
			// go func ()  {
			// 	err1 =getDveNexusData(TOKEN, dveURL)
			// 	wg.Done()
			// }()
			wg.Wait()

			if err1 != nil {
				os.Exit(12)
			}

			if is, _ := cn.IsSn(sn); is{
				return true
			}
		}
	case "watsons":
		{
			var  err1  error
			wg.Add(num)
			go func ()  {
				err1 = getCpeWatsonsData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func ()  {
				err1 = getPopWatsonsData(TOKEN, popURL)
				wg.Done()
			}()
			// go func ()  {
			// 	err1 =getDveWatsonsData(TOKEN, dveURL)
			// 	wg.Done()
			// }()
			wg.Wait()

			if err1 != nil {
				os.Exit(12)
			}

			if is, _ := cw.IsSn(sn); is{
				return true
			}
		}
	case "watsonsha":
		{
			var  err1  error
			wg.Add(num)
			go func ()  {
				err1 = getCpeWatsonsHaData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func ()  {
				err1 = getPopWatsonsHaData(TOKEN, popURL)
				wg.Done()
			}()
			// go func ()  {
			// 	err1 =getDveWatsonsHaData(TOKEN, dveURL)
			// 	wg.Done()
			// }()
			wg.Wait()

			if err1 != nil {
				os.Exit(12)
			}

			if is, _ := ch.IsSn(sn); is{
				return true
			}
		}
	case "tassadar":
		{
			var  err1  error
			wg.Add(num)
			go func ()  {
				err1 = getCpeZeratulData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func ()  {
				err1 = getPopZeratulData(TOKEN, popURL)
				wg.Done()
			}()
			// go func ()  {
			// 	err1 =getDveZeratulData(TOKEN, dveURL)
			// 	wg.Done()
			// }()
			wg.Wait()

			if err1 != nil {
				os.Exit(12)
			}
			if is, _ := cz.IsSn(sn); is{
				return true
			}
		}
	}
	return false
}

// 根据sn去云端接口获取属于哪个mode sdwan6
func GetModebySevenSn(sn string) string {
	if err = getOperationData(TOKEN, neter.GetOperationFromRoute()); err != nil {
		os.Exit(11)
	}
	return op.SnInMode(sn)
}

// 企业号使用 已知mode获取cpe,dve,pop数据并放入到内存 支持nexus/valor/tassadar
func SyncEnDataMemorybyMode(mode string) {
	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	dveURL := neter.GetDveFromRoute(mode)
	switch mode {
	case "valor":
		{
			err = getCpeValorData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopValorData(TOKEN, popURL)
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
			err = getCpeZeratulData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopZeratulData(TOKEN, popURL)
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
			err = getCpeNexusData(TOKEN, cpeURL)
			if err != nil {
				os.Exit(12)
			}
			err = getPopNexusData(TOKEN, popURL)
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