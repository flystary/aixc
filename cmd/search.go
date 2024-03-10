package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("write", "w", "", "Write current data to a file")
}

var searchCmd = &cobra.Command{
	Use:   "search [<SN>, <SN>,....]",
	Short: "Print multiple or single rows of data in table form",
	Long:  `Use search to display single or multiple ucpe data`,
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		writeOpt, err := cmd.Flags().GetString("write")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		var aixc Cmd = &CLI{
			sns:   args,
			Write: cli.DecodeWrite(writeOpt),
		}
		aixc.run()
	},
}
