package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("Mseven", "m", false, "If ucpe belongs to 6.X platform, Please use it")
}

var listCmd = &cobra.Command{
	Use:   	 "list [<SN>, <SN>,....]",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Print multiple rows of data in tabular form",
	Long: 	 `use list to list everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		mseven, err := cmd.Flags().GetBool("Mseven")
		if err != nil {
			println("getBool err: ", err)
			return
		}
		if mseven {
			showSevenMany(args)
		} else {
			showMany(args)
		}
	},
}