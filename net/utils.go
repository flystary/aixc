package net

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"

	"aixc/model"

	"github.com/olekukonko/tablewriter"
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

func timeUnix(e time.Time) int64 {
	return e.UnixNano() / 1e6
}

func newMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
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
		mode := model.M(i).Enum()
		go getMapByChan(sn, mode, limit, wg)
	}
	wg.Wait()

	close(channel)
	return useMapArrgetMode(arr)
}

func getMapByChan(sn, mode string, limit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	relationMap := make(map[string]string, 6)
	is  := "No"
	if syncDataMemorybySnMode(sn, mode) {
		is = "Yes"
	}
	relationMap[mode] = is
	select {
		case channel <- relationMap:
			time.Sleep(1*time.Millisecond)
		default:
			close(channel)
	}
	<- limit
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// HOME 用户根路径
func HOME() string {
	dir, _ := homedir.Dir()
	return dir
}