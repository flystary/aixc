package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("mode", "m", "","Appoint the UCPE Mode")
	listCmd.Flags().BoolP("seven", "", false, "Appoint that the ucpe Mode belongs to SDWAN6")
}

var listCmd = &cobra.Command{
	Use:   	 "list [<SN>, <SN>,....]",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Print multiple or single rows of data in table form",
	Long: 	 `Use list to display single or multiple ucpe data`,
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
					showManybyModeSns(mode,args)
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