package net

import (
	"sync"
	"aixc/model"
)

func modeController(simc map[string]string ) {
	arr = append(arr, simc)
}

func ThreadQueryMode(sn string) string {
	wg := &sync.WaitGroup{}
	for i := 1; i < 6; i++ {
		wg.Add(1)
		mode := model.M(i).Enum()
		go getMapByChan(sn, mode, wg)
	}
	wg.Wait()

	return useMapArrgetMode(arr)
}

func getMapByChan(sn, mode string, wg *sync.WaitGroup) {

	relationMap := make(map[string]string, 1)
	defer wg.Done()
	is := "No"
	// now := time.Now()
	if SyncDataMemorybySnMode(sn, mode) {
		is = "Yes"
	}
	// useTime := time.Since(now)
	// fmt.Println(mode ,"use time", useTime)

	relationMap[mode] = is
	modeController(relationMap)
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
				m = i
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