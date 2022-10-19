package cmd

import (
	"github.com/spf13/cobra"

	"aixc/net"
)

func init() {}

var rootCmd = &cobra.Command{
	Use: "aixc",
}

func show(sn string) {
	net.Search(sn)
}

func showSeven(sn string) {
	net.SearchSeven(sn)
}

func showMany(sns []string) {
	net.SearchMany(sns)
}

func showSevenMany(sns []string) {
	net.SearchSevenMany(sns)
}

func showManybyMode(mode string){
	println("此命令暂不支持!")
}
// Run cmd
func Run() {
	rootCmd.Execute()
}