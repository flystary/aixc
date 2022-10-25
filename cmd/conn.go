package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(connCmd)
	connCmd.Flags().BoolP("seven", "s", false, "If ucpe belongs to SDWAN6 platform, Please use it")
}

var connCmd = &cobra.Command{
	Use:   	 "conn <SN>",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Use this function to connect CPE or node",
	Long: 	 `use conn to connet everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(_ *cobra.Command, args []string) {
		show(args[0])
	},
}