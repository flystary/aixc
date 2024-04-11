package cmd

import (
	. "aixc/tools/color"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Model      = Cyan("model")
	Version    = Cyan("version")
	Pop        = Cyan("pop")
	Enterprise = Cyan("enterprise")

	selectOption = fmt.Sprintf("%s|%s|%s|%s", Model, Version, Pop, Enterprise)
)

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("mode", "m", "valor", fmt.Sprintf("Appoint the UCPE Mode, Option is %s", modeOption))
	showCmd.Flags().StringP("select", "s", "", fmt.Sprintf("Appoint the filter object of UCPE, Option is %s", selectOption))
	showCmd.Flags().StringP("enterprise", "e", "", "Appoint that the filtering object of UCPE is the enterprise number")
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Print your filteCyan data in tabular form",
	Long:  `Use show to list the UCPE of the specified options according to your filter`,
	// Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		entn, err := cmd.Flags().GetString("enterprise")
		if err != nil {
			println("getbool err: ", err)
			return
		}

		option, err := cmd.Flags().GetString("select")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		if checkShowMode(mode) {
			if len(entn) == 0 && len(option) == 0 {
				if mode == "tassadar" || mode == "watsonsha" || mode == "yifeng" {
					showMode(mode)
					return
				} else {
					println("不支持此用法!")
					return
				}

			} else if len(entn) > 0 && len(option) > 0 {
				println("不支持此用法!")
				return
			} else {
				if len(entn) > 0 && len(option) == 0 {
					Entn(mode, entn)
					return
				} else {
					Option(mode, option, args[0])
					return
				}
			}
		}
	},
}

func checkShowMode(mode string) bool {
	if mode == "valor" || mode == "yifeng" || mode == "tassadar" || mode == "watsonsha" || mode == "watsons" {
		return true
	}
	return false
}

func Option(mode, option, args string) {
	switch option {
	case "model":
		showModel(mode, args)
		return
	case "version":
		showVersion(mode, args)
		return
	case "pop":
		showPop(mode, args)
		return
	case "enterprise":
		showEnterprise(mode, args)
		return
	default:
		println("不支持此用法!")
		return
	}
}

func Entn(mode, entn string) {
	if mode == "watsons" {
		println("不支持此用法!")
		return
	}
	showEnterprise(mode, entn)
}
