package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of xc",
	Long:  `All software has versions. This is aixc's`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Aixc Version: v2.0 -- HEAD")
	},
}