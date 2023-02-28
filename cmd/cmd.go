package cmd

import (
	log "aixc/utils/log"
	"fmt"

	n "aixc/net"

	"github.com/spf13/cobra"
)

func init() {}

var rootCmd = &cobra.Command{
	Use: "aixc",
}

var (
	Valor      = n.Cyan("valor")
	Tassadar   = n.Cyan("tassadar")
	Watsons    = n.Cyan("watsons")
	WatsonsHa  = n.Cyan("watsonsha")
	Nexus      = n.Cyan("nexus")
	modeOption = fmt.Sprintf("%s|%s|%s|%s|%s", Valor, Tassadar, Watsons, WatsonsHa, Nexus)
)

// func show(sn string) {
// 	n.SearchBySn(sn)
// }

func showSeven(sn string) {
	log.Infof("showSeven sn:%s", sn)
	n.SearchSevenBySn(sn)
}

func showMany(sns []string) {
	log.Infof("showMany sns:%s", sns)
	n.SearchManyBySns(sns)
}

func showSevenMany(sns []string) {
	log.Infof("showSevenMany sns:%s", sns)
	n.SearchSevenManyBySns(sns)
}

func showManybyModeSns(mode string, sns []string) {
	log.Infof("showManybyModeSns mode:%s sns:%s", mode, sns)
	n.SearchByModeSns(mode, sns)
}

func showMode(mode string) {
	log.Infof("showMode mode:%s", mode)
	n.SearchByMode(mode)
}

// filter

func showModel(mode, model string) {
	log.Infof("showModel mode:%s model:%s", mode, model)
	n.FilterModelByMode(mode, model)
}

func showVersion(mode, version string) {
	log.Infof("showVersion mode:%s version:%s", mode, version)
	n.FilterVersionByMode(mode, version)
}

func showPop(mode, addr string) {
	log.Infof("showPop mode:%s addr:%s", mode, addr)
	n.FilterPopByMode(mode, addr)
}

func showEnterprise(mode, name string) {
	log.Infof("showEnterprise mode:%s Enterprise:%s", mode, name)
	n.FilterEnterpriseByMode(mode, name)
}

// Run cmd
func Run() {
	rootCmd.Execute()
}
