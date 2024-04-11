package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search [<SN>, <SN>,....]",
	Short: "Print multiple or single rows of data in table form",
	Long:  `Use search to display single or multiple ucpe data`,
	Args:  cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

	    if len(args) >= 1 {
            showMany(args)
	    } else {
	        println("不支持此用法!")
	        return
        }
    },
}

