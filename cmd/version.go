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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Aixc Version: v1.0 -- HEAD")
	},
}