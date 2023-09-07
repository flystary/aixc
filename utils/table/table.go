package table


import (
	"os"
	"github.com/olekukonko/tablewriter"
)


func TableBasic(data [][]string) {
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

func Table2Basic(data [][]string) {
	ALIGN_LEFT := 3
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ucpesn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip", "port", "enterprise", "alias"})
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


func Table3Basic(data [][]string) {
	ALIGN_LEFT := 3
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ucpesn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip",  "port", "alias"})
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

func TableMarkdown(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"sn", "model", "version", "updatetime", "masterpopip", "mastercpeip", "backuppopip", "backupcpeip"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
}
