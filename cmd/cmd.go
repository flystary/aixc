package cmd

import (
	"github.com/spf13/cobra"
	n "aixc/net"
)

func init() {}

var rootCmd = &cobra.Command{
	Use: "aixc",
}

func show(sn string) {
	n.Search(sn)
}

func showSeven(sn string) {
	n.SearchSeven(sn)
}

func showMany(sns []string) {
	n.SearchMany(sns)
}

func showSevenMany(sns []string) {
	n.SearchSevenMany(sns)
}

func showManybyMode(mode string){
	println("此命令暂不支持!", mode)
}
// Run cmd
func Run() {
	rootCmd.Execute()
}