package net

import (
	"fmt"
	"sync"
	"time"
)


var (
	arr  []map[string]string
	snInModeChannel = make(chan map[string]string, 15)
)
//Mtype 自定义int
type Mtype int

const (
	valor Mtype = iota + 1
	nexus
	watsons
	watsonsHa
	tassadar
)

func init() {
	go snInModeController()
	arr = make([]map[string]string, 1)
	arr[0] = make(map[string]string)

}

func (m Mtype) enum() string {
	var mode string
	switch m {
	case valor:
		{
			mode = "valor"
		}
	case nexus:
		{
			mode = "nexus"
		}
	case watsons:
		{
			mode = "watsons"
		}
	case watsonsHa:
		{
			mode = "watsons_ha"
		}
	case tassadar:
		{
			mode = "tassadar"
		}

	default:
		{
			mode = "unknown"
		}
	}
	return mode
}

func snInMode(sn, mode string, limit chan bool, wg *sync.WaitGroup) {
	// var snInModeMap map[string]string
	defer wg.Done()
	snInModeMap := make(map[string]string, 6)
	is  := "No"
	if isSnInMode(sn, mode) {
		is = "Yes"
	}
	snInModeMap[mode] = is

	select {
		case snInModeChannel <- snInModeMap: {
			time.Sleep(1*time.Second)
		}
		default:
			fmt.Println("通道已满", len(snInModeChannel))
			close(snInModeChannel)
	}

	<- limit
}

func snInModeController() {
	for simc := range snInModeChannel {
		arr = append(arr, simc)
	}
}

func getSnInMode(sn string) string {
	wg := &sync.WaitGroup{}
	limit := make(chan bool, 20)
	for i := 1; i < 6; i++ {
		wg.Add(1)
		limit <- true
		mode := Mtype(i).enum()
		go snInMode(sn, mode, limit, wg)
	}
	wg.Wait()

	close(snInModeChannel)
	return useMapArrgetMode(arr)
}
