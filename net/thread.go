package net

import (
	"sync"

	"aixc/model"
)

var (
	arr []map[string]string
)


func init() {
    arr = make([]map[string]string, 0, 6)

}

func modeController(simc map[string]string ) {
	arr = append(arr, simc)
}

func ThreadQueryMode(sn string) string {
	wg := &sync.WaitGroup{}
	for i := 1; i < 5; i++ {
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
	ms := make([]string, 0)
	var m string
	for _, mp := range marr {
		for k, v := range mp {
			if v == "Yes" && k != "" {
				ms = append(ms, k)
				break
			}
		}
	}

	if len(ms) >= 2 {
		for _, i := range ms {
			m = "unknown"
            if i == "yifeng" {
                m = i
                break
            }
			if i == "valor" {
				m = i
				break
			}
		}
	} else if len(ms) == 1 {
		m = ms[0]
	} else {
		m = "unknown"
	}
	return m
}
