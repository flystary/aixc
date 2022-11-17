package net

import (
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"

	"aixc/model"
)

var (
	arr  []map[string]string
)

func init() {
	arr = make([]map[string]string, 0, 6)
}

func Red(iput string) string {
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	return red(iput)
}

func Green(iput string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	return green(iput)
}

func Cyan(iput string) string {
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	return cyan(iput)
}

func Blue(iput string) string {
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	return blue(iput)
}

func White(iput string) string {
	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	return white(iput)
}

func tableBasic(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func tableMarkdown(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
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

// HOME 用户根路径
func HOME() string {
	dir, _ := homedir.Dir()
	return dir
}