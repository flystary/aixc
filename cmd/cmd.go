package cmd

import (
	. "aixc/net"
	. "aixc/tools/color"
	xlog "aixc/tools/log"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {}

var rootCmd = &cobra.Command{
	Use: "aixc",
}

var (
	Valor      = Cyan("valor")
	Yifeng     = Cyan("yifeng")
	Tassadar   = Cyan("tassadar")
	Watsons    = Cyan("watsons")
	WatsonsHa  = Cyan("watsonsha")
	modeOption = fmt.Sprintf("%s|%s|%s|%s|%s", Valor, Tassadar, Yifeng, Watsons, WatsonsHa)
)

// func show(sn string) {
// 	n.SearchBySn(sn)
// }

func showSeven(sn string) {
	xlog.Infof("showSeven sn:%s", sn)
	SearchSevenBySn(sn)
}

func showMany(sns []string) {
	xlog.Infof("showMany sns:%s", sns)
	SearchManyBySns(sns)
}

func showSevenMany(sns []string) {
	xlog.Infof("showSevenMany sns:%s", sns)
	SearchSevenManyBySns(sns)
}

func showManybyModeSns(mode string, sns []string) {
	xlog.Infof("showManybyModeSns mode:%s sns:%s", mode, sns)
	SearchByModeSns(mode, sns)
}

func showMode(mode string) {
	xlog.Infof("showMode mode:%s", mode)
	SearchByMode(mode)
}

// filter

func showModel(mode, model string) {
	xlog.Infof("showModel mode:%s model:%s", mode, model)
	FilterModelByMode(mode, model)
}

func showVersion(mode, version string) {
	xlog.Infof("showVersion mode:%s version:%s", mode, version)
	FilterVersionByMode(mode, version)
}

func showPop(mode, addr string) {
	xlog.Infof("showPop mode:%s addr:%s", mode, addr)
	FilterPopByMode(mode, addr)
}

func showEnterprise(mode, name string) {
	xlog.Infof("showEnterprise mode:%s Enterprise:%s", mode, name)
	FilterEnterpriseByMode(mode, name)
}

// Run cmd
func Run() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.Execute()
}
