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

func showEnterprise(mode, enterprise string) {
	n.SearchByModeEnterprise(mode, enterprise)
}

func showManybyModeSns(mode string, sns []string){
	n.SearchByModeSns(mode,sns)
}
// Run cmd
func Run() {
	rootCmd.Execute()
}