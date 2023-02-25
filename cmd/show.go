package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("select", "s", "","Appoint the filter object of UCPE")
	showCmd.Flags().StringP("mode", "m", "","Appoint the UCPE Mode")
	showCmd.Flags().StringP("enterprise", "e", "null","Appoint that the filtering object of UCPE is the enterprise number")

}

var showCmd = &cobra.Command{
	Use:   	 "show",
	Example: "xc show -m <[MODE]> -s <[OPTION]>",
	Short:   "Print your filtered data in tabular form",
	Long: 	 `Use show to list the UCPE of the specified options according to your filter`,
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

		if mode == "valor" || mode == "nexus" || mode == "tassadar" || mode == "watsonsha" || mode == "" {

			if (mode == "tassadar" && entn == "null")  || (mode == "watsonsha" && entn == "null")  {
				showMode(mode)
				return
			}

			if mode == "watsonsha" {
				return
			}

			if mode == "" {
				mode = "valor"
			}

			switch option {
			case "model":
				showModel(mode, args[0])
			case "version":
				showVersion(mode, args[0])
			case "pop":
				showPop(mode, args[0])
			case "enterprise":
				showEnterprise(mode, args[0])
			default:
				return
			}

			if entn == "null" {
				println("enterprise is Error")
				return
			}
			showEnterprise(mode, entn)
		}
	},

}