package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
	// showCmd.Flags().StringP("select", "s", "","Appoint the filter object of UCPE")
	showCmd.Flags().StringP("mode", "m", "","Appoint the UCPE Mode")
	showCmd.Flags().StringP("enterprise", "e", "","Appoint that the filtering object of UCPE is the enterprise number")

}

var showCmd = &cobra.Command{
	Use:   	 "show",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Print your filtered data in tabular form",
	Long: 	 `Use show to list the UCPE of the specified options according to your filter`,
	// Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, _ []string) {

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
		// stype, err := cmd.Flags().GetString("select")
		// if err != nil {
		// 	println("getstring err: ", err)
		// 	return
		// }

		if mode == "valor" || mode == "nexus" || mode == "tassadar" || mode == "" {
			if mode == "" {
				mode = "valor"
			}
			showEnterprise(mode, entn)
		}
	},
}
