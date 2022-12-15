package net

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
)

var (
	arr []map[string]string
)

func init() {
	arr = make([]map[string]string, 0, 6)
}

// HOME 用户根路径
func HOME() string {
	dir, _ := homedir.Dir()
	return dir
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
	ALIGN_LEFT := 3
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ucpesn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(ALIGN_LEFT)
	table.SetAlignment(ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func table2Basic(data [][]string) {
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

func Put(URL string, Payload *strings.Reader)  *http.Response {

	req, _ := http.NewRequest("PUT", URL, Payload)

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	return resp
}
