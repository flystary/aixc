package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolP("seven", "", false, "If ucpe belongs to SDWAN6 platform, Please use it")
}

var showCmd = &cobra.Command{
	Use:     "show <SN>",
	// Example: "xc show 7x00xxxxxxxxxxx",
	Short:   "Print single line data in tabular form",
	Long:    `use show to show everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		mseven, err := cmd.Flags().GetBool("seven")
		if err != nil {
			println("getBool err: ", err)
			return
		}
		if mseven {
			showSeven(args[0])
		} else {
			show(args[0])
		}
	},
}