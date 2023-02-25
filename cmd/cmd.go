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

func showManybyModeSns(mode string, sns []string){
	n.SearchByModeSns(mode,sns)
}

func showMode(mode string) {
	n.SearchByMode(mode)
}

// filter

func showModel(mode, model string) {
	n.FilterModelByMode(mode, model)
}

func showVersion(mode, version string) {
	n.FilterVersionByMode(mode, version)
}

func showPop(mode, addr string) {
	n.FilterPopByMode(mode, addr)
}

func showEnterprise(mode, name string) {
	n.FilterEnterpriseByMode(mode, name)
}

// Run cmd
func Run() {
	rootCmd.Execute()
}