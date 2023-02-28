package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("mode", "m", "", fmt.Sprintf("Appoint the UCPE Mode, Option is %s", modeOption))
	listCmd.Flags().BoolP("seven", "", false, "Appoint that the ucpe Mode belongs to SDWAN6")
}

var listCmd = &cobra.Command{
	Use:   "list [<SN>, <SN>,....]",
	Short: "Print multiple or single rows of data in table form",
	Long:  `Use list to display single or multiple ucpe data`,
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		useSeven, err := cmd.Flags().GetBool("seven")
		if err != nil {
			println("getBool err: ", err)
			return
		}
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		if checkListMode(mode) {
			if len(mode) == 0 {
				if useSeven {
					if len(args) == 1 {
						showSeven(args[0])
					} else {
						showSevenMany(args)
					}
				} else {
					showMany(args)
					return
				}
			} else {
				if useSeven {
					println("不支持此用法!")
					return
				}
				showManybyModeSns(mode, args)
				return
			}
		} else {
			println("不支持此用法!")
			return
		}
	},
}

func checkListMode(mode string) bool {
	if mode == "valor" || mode == "nexus" || mode == "watsons" || mode == "watsonsha" || mode == "tassadar" || mode == "" {
		return true
	}
	return false
}

func Seven(is bool) {

}
