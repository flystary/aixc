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
	n.SearchBySn(sn)
}

func showSeven(sn string) {
	n.SearchSevenBySn(sn)
}

func showMany(sns []string) {
	n.SearchManyBySns(sns)
}

func showSevenMany(sns []string) {
	n.SearchSevenManyBySns(sns)
}

func showEnterprise(mode,enterprise string) {
	n.SearchByModeEnterprise(mode, enterprise)
}

func showManybyMode(mode string){
	println("此命令暂不支持!", mode)
}
// Run cmd
func Run() {
	rootCmd.Execute()
}