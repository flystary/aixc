package control


import (
	"os"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func red(iput string) string {
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	return red(iput)
}

func green(iput string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	return green(iput)
}

func cyan(iput string) string {
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	return cyan(iput)
}

func blue(iput string) string {
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	return blue(iput)
}

func white(iput string) string {
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