package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   	 "exec <SN>",
	Short:   "Use this option to remotely execute commands on UCPE or POP",
	Long: 	 `This method allows CPE or POP to execute commands`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(_ *cobra.Command, args []string) {
		println("不支持此用法!")
		return
	},
}
