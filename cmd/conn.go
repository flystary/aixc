package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(connCmd)
}

var connCmd = &cobra.Command{
	Use:   	 "conn <SN>",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Use this option to connect CPE or POP remotely",
	Long: 	 `Use conn to remotely connect the UCPE and POP you need`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		println("不支持此用法!")
		// return
	},
}