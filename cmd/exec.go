package cmd

import (
	"aixc/net"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   	 "exec <SN>",
	Short:   "Use this method to execute commands",
	Long: 	 `This method allows CPE or POP to execute commands`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(_ *cobra.Command, args []string) {
		net.Search(args[0])
	},
}