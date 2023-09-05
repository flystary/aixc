package net

import (
	"fmt"
	"os"
	"sync"

	r "aixc/route"
)

var (
	// TOKEN
	TOKEN string
	err   error
	route r.Route
	neter Neter = &route
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
type Requester interface {
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
	// path := fmt.Sprintf("./%s", fileName)
	route = r.LoadRoute(path)
	// 获取token
	TOKEN = GetToken(neter.GetTokenFromRoute())
}

// 已知mode获取cpe,dve,pop数据并放入到内存
func SyncDataMemorybyMode(mode string) {
	wg := &sync.WaitGroup{}
	num := 3

	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	dveURL := neter.GetDveFromRoute(mode)
	switch mode {
	case "valor":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeValorData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopValorData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveValorData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	case "tassadar":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeZeratulData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopZeratulData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveZeratulData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	case "watsons":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeWatsonsData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopWatsonsData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveWatsonsData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	case "watsonsha":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeWatsonsHaData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopWatsonsHaData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveWatsonsHaData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	case "nexus":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeNexusData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopNexusData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveNexusData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	}
}

// 已知sn和随机mode获取cpe,dve,pop数据并放入到内存
func SyncDataMemorybySnMode(sn, mode string) bool {
	wg := &sync.WaitGroup{}
	num := 2
	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	// dveURL := neter.GetDveFromRoute(mode)
	switch mode {
	case "valor":
		{
			var cpeErr error
			var popErr error
			// var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeValorData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopValorData(TOKEN, popURL)
				wg.Done()
			}()
			// go func() {
			// 	dveErr = getDveValorData(TOKEN, dveURL)
			// 	wg.Done()
			// }()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			// if dveErr != nil {
			// 	os.Exit(15)
			// }

			if is, _ := cv.IsSn(sn); is {
				return true
			}
		}
	case "nexus":
		{
			var cpeErr error
			var popErr error
			// var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeNexusData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopNexusData(TOKEN, popURL)
				wg.Done()
			}()
			// go func() {
			// 	dveErr = getDveNexusData(TOKEN, dveURL)
			// 	wg.Done()
			// }()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			// if dveErr != nil {
			// 	os.Exit(15)
			// }

			if is, _ := cn.IsSn(sn); is {
				return true
			}
		}
	case "watsons":
		{

			var cpeErr error
			var popErr error
			// var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeWatsonsData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopWatsonsData(TOKEN, popURL)
				wg.Done()
			}()
			// go func() {
			// 	dveErr = getDveWatsonsData(TOKEN, dveURL)
			// 	wg.Done()
			// }()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			// if dveErr != nil {
			// 	os.Exit(15)
			// }

			if is, _ := cw.IsSn(sn); is {
				return true
			}
		}
	case "watsonsha":
		{
			var cpeErr error
			var popErr error
			// var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeWatsonsHaData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopWatsonsHaData(TOKEN, popURL)
				wg.Done()
			}()
			// go func() {
			// 	dveErr = getDveWatsonsHaData(TOKEN, dveURL)
			// 	wg.Done()
			// }()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			// if dveErr != nil {
			// 	os.Exit(15)
			// }

			if is, _ := ch.IsSn(sn); is {
				return true
			}
		}
	case "tassadar":
		{
			var cpeErr error
			var popErr error
			// var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeZeratulData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopZeratulData(TOKEN, popURL)
				wg.Done()
			}()
			// go func() {
			// 	dveErr = getDveNexusData(TOKEN, dveURL)
			// 	wg.Done()
			// }()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			// if dveErr != nil {
			// 	os.Exit(15)
			// }

			if is, _ := cz.IsSn(sn); is {
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
	wg := &sync.WaitGroup{}
	num := 3

	cpeURL := neter.GetCpeFromRoute(mode)
	popURL := neter.GetPopFromRoute(mode)
	dveURL := neter.GetDveFromRoute(mode)
	switch mode {
	case "valor":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeValorData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopValorData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveValorData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	case "tassadar":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeZeratulData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopZeratulData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveZeratulData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	case "nexus":
		{
			var cpeErr error
			var popErr error
			var dveErr error

			wg.Add(num)
			go func() {
				cpeErr = getCpeNexusData(TOKEN, cpeURL)
				wg.Done()
			}()
			go func() {
				popErr = getPopNexusData(TOKEN, popURL)
				wg.Done()
			}()
			go func() {
				dveErr = getDveNexusData(TOKEN, dveURL)
				wg.Done()
			}()

			wg.Wait()

			if cpeErr != nil {
				os.Exit(12)
			}
			if popErr != nil {
				os.Exit(13)
			}
			if dveErr != nil {
				os.Exit(15)
			}
		}
	}
}
