package net

import (
	"os"
	"fmt"
	"time"
	"sync"
)

var (
	arr  []map[string]string
	channel = make(chan map[string]string, 15)
)

func init() {
	go modeController()
	arr = make([]map[string]string, 1)
	arr[0] = make(map[string]string)

}

func timeUnix(e time.Time) int64 {
	return e.UnixNano() / 1e6
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

func modeController() {
	for simc := range channel {
		arr = append(arr, simc)
	}
}

func threadQueryMode(sn string) string {
	wg := &sync.WaitGroup{}
	limit := make(chan bool, 20)
	for i := 1; i < 6; i++ {
		wg.Add(1)
		limit <- true
		mode := Mtype(i).enum()
		go getMapByChan(sn, mode, limit, wg)
	}
	wg.Wait()

	close(channel)
	return useMapArrgetMode(arr)
}

func getMapByChan(sn, mode string, limit chan bool, wg *sync.WaitGroup) {
	// var relationMap map[string]string
	defer wg.Done()
	relationMap := make(map[string]string, 6)
	is  := "No"
	if syncDataMemorybySnMode(sn, mode) {
		is = "Yes"
	}
	relationMap[mode] = is
	select {
		case channel <- relationMap: {
			time.Sleep(1*time.Second)
		}
		default:
			fmt.Println("通道已满", len(channel))
			close(channel)
	}
	<- limit
}

func snInSevenMode(sn string) string {
	if opo, err = getOperationData(token, rules.OperationRouteByMode()); err != nil {
		os.Exit(11)
	}
	return opo.SnInMode(sn)
}