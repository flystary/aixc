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
		var ops *Options

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
		ops = ops.isEnterprise(entn)

		sele, err := cmd.Flags().GetString("select")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		var wr *Write
		write, err := cmd.Flags().GetString("write")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		var aixc Cmd = &CLI{
			mode:    mode,
			options: ops.Select(sele, args),
			write:   *wr.Decode(write),
		}
		aixc.run()
	},
}
