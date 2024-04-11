package net

import (
	"aixc/model"
	"sync"
)

func syncCpe() {
	wg := &sync.WaitGroup{}
	num := 5
	wg.Add(num)

	go func() {
		SyncAllDataMemory("valor")
		wg.Done()
	}()
	go func() {
		SyncAllDataMemory("yifeng")
		wg.Done()
	}()
	go func() {
		SyncAllDataMemory("tassadar")
		wg.Done()
	}()
	go func() {
		SyncAllDataMemory("watsons")
		wg.Done()
	}()
	go func() {
		SyncAllDataMemory("watsonsha")
		wg.Done()
	}()
	wg.Wait()
}

func SearchSnByMode(sn string) string {
	var m string
	ms := make([]string, 0)
	modeExs := make([]map[string]bool, 0, 6)
	for i := 1; i < 5; i++ {
		mode := model.M(i).Enum()
		modeEx := make(map[string]bool, 1)
		switch mode {
		case "valor":
			{
				modeEx[mode] = cv.IsExist(sn)
				modeExs = append(modeExs, modeEx)
			}
		case "yifeng":
			{
				modeEx[mode] = cy.IsExist(sn)
				modeExs = append(modeExs, modeEx)
			}
		case "tassadar":
			{
				modeEx[mode] = cz.IsExist(sn)
				modeExs = append(modeExs, modeEx)
			}
		case "watsons":
			{
				modeEx[mode] = cw.IsExist(sn)
				modeExs = append(modeExs, modeEx)
			}
		case "watsonsha":
			{
				modeEx[mode] = ch.IsExist(sn)
				modeExs = append(modeExs, modeEx)
			}
		}
	}

	for _, mEs := range modeExs {
		for k, v := range mEs {
			if v {
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
