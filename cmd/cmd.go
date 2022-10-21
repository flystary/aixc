package cmd

import (
	"github.com/spf13/cobra"

	m "aixc/manage"
)

func init() {}

var rootCmd = &cobra.Command{
	Use: "aixc",
}

func show(sn string) {
	m.Search(sn)
}

func showSeven(sn string) {
	m.SearchSeven(sn)
}

func showMany(sns []string) {
	m.SearchMany(sns)
}

func showSevenMany(sns []string) {
	m.SearchSevenMany(sns)
}

func showManybyMode(mode string){
	println("此命令暂不支持!")
}
// Run cmd
func Run() {
	rootCmd.Execute()
}