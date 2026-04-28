package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func newTable() *tablewriter.Table {
	t := tablewriter.NewWriter(os.Stdout)

	t.SetAutoWrapText(false)
	t.SetAutoFormatHeaders(true)

	t.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	t.SetAlignment(tablewriter.ALIGN_LEFT)

	t.SetCenterSeparator("")
	t.SetColumnSeparator("")
	t.SetRowSeparator("")

	t.SetHeaderLine(false)
	t.SetBorder(false)

	t.SetTablePadding("\t")
	t.SetNoWhiteSpace(true)

	return t
}

func render(headers []string, data [][]string) {
	t := newTable()
	t.SetHeader(headers)
	t.AppendBulk(data)
	t.Render()
}

func TableBasic(data [][]string) {
	render([]string{
		"ucpesn", "model", "version", "updatetime",
		"masterpopip", "mastercpeip",
		"backuppopip", "backupcpeip",
	}, data)
}

func Table2Basic(data [][]string) {
	render([]string{
		"ucpesn", "model", "version", "updatetime",
		"masterpopip", "mastercpeip",
		"backuppopip", "backupcpeip",
		"port", "enterprise", "alias",
	}, data)
}

func Table3Basic(data [][]string) {
	render([]string{
		"ucpesn", "model", "version", "updatetime",
		"masterpopip", "mastercpeip",
		"backuppopip", "backupcpeip",
		"port", "alias",
	}, data)
}

func TableMarkdown(data [][]string) {
	t := tablewriter.NewWriter(os.Stdout)

	t.SetHeader([]string{
		"sn", "model", "version", "updatetime",
		"masterpopip", "mastercpeip",
		"backuppopip", "backupcpeip",
	})

	t.SetBorders(tablewriter.Border{
		Left: true, Right: true,
	})

	t.SetCenterSeparator("|")

	t.AppendBulk(data)
	t.Render()
}
