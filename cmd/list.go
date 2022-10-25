package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("mode", "m", "","ucpe mode, Please use it")
	listCmd.Flags().BoolP("seven", "s", false, "if ucpe belongs to SDWAN6 platform, Please use it")
}

var listCmd = &cobra.Command{
	Use:   	 "list [<SN>, <SN>,....]",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Print multiple rows of data in tabular form",
	Long: 	 `use list to list everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		isSeven, err := cmd.Flags().GetBool("seven")
		if err != nil {
			println("getBool err: ", err)
			return
		}
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		if mode == "valor" || mode == "nexus" || mode == "watsons" || mode == "watsonsha" || mode == "tassadar" || mode == "" {
			// isSeven
			if isSeven && mode == "" {
				showSevenMany(args)
			}else if isSeven && mode != ""{
				println("不支持此用法!")
				return
			}else{
				if mode != "" {
					showManybyMode(args[0])
					return
				}
				showMany(args)
			}
		}else{
			println("不支持此用法!")
			return
		}
	},
}