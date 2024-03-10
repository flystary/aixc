package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("mode", "m", "valor", fmt.Sprintf("Appoint the UCPE Mode, Option is %s", modeOption))
	showCmd.Flags().StringP("select", "s", "", fmt.Sprintf("Appoint the filter object of UCPE, Option is %s", selectOption))
	showCmd.Flags().StringP("enterprise", "e", "", "Appoint that the filtering object of UCPE is the enterprise number")
	showCmd.Flags().StringP("write", "w", "", "Write current data to a file")
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Print your filteCyan data in tabular form",
	Long:  `Use show to list the UCPE of the specified options according to your filter`,
	// Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		modeopt, err := cmd.Flags().GetString("mode")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		enterpriseopt, err := cmd.Flags().GetString("enterprise")
		if err != nil {
			println("getbool err: ", err)
			return
		}
		if enterpriseopt != "" {
			cli.isEnterprise = true
			cli.enterprise = enterpriseopt
		}

		selectopt, err := cmd.Flags().GetString("select")
		if err != nil {
			println("getstring err: ", err)
			return
		}
		writeOpt, err := cmd.Flags().GetString("write")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		var aixc Cmd = &CLI{
			mode: modeopt,
			// enterprise: entn,
			// ops.isEnte = true
			Options: cli.SelectOpt(selectopt, args),
			Write:   cli.DecodeWrite(writeOpt),
		}
		aixc.run()
	},
}
