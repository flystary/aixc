package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("mode", "m", "", fmt.Sprintf("Appoint the UCPE Mode, Option is %s", modeOption))
	listCmd.Flags().BoolP("seven", "", false, "Appoint that the ucpe Mode belongs to SDWAN6")
	listCmd.Flags().StringP("write", "w", "", "Write current data to a file")
}

var listCmd = &cobra.Command{
	Use:   "list [<SN>, <SN>,....]",
	Short: "Print multiple or single rows of data in table form",
	Long:  `Use list to display single or multiple ucpe data`,
	Args:  cobra.MinimumNArgs(1),

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

		var wr	*Write
		write, err := cmd.Flags().GetString("write")
		if err != nil {
			println("getstring err: ", err)
			return
		}
		var aixc Cmd = &CLI{
			sns: args,
			mode: mode,
			isSeven: isSeven,
			write: *wr.Decode(write),
		}
		aixc.run()
	},
}
