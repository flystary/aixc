package cmd

import (
	"fmt"
	"aixc/net"
	"github.com/spf13/cobra"
)


// RootCmd root
var RootCmd = &cobra.Command{
	Use: "aixc",
}

var cmdShow = &cobra.Command{
	Use:     "show <SN>",
	// Example: "xc show 7x00114401917b5f0",
	Short:   "Print single line data in tabular form",
	Long:    `use show to show everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		mseven, err := cmd.Flags().GetBool("Mseven")
		if err != nil {
			fmt.Println("getBool err: ", err)
			return
		}
		if mseven {
			showSeven(args[0])
		} else {
			show(args[0])
		}
	},
}

var cmdConn = &cobra.Command{
	Use:   	 "conn <SN>",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Use this function to connect CPE or node",
	Long: 	 `use conn to connet everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(_ *cobra.Command, args []string) {
		net.Search(args[0])
	},
}

var cmdList = &cobra.Command{
	Use:   	 "list [<SN>, <SN>,....]",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Print multiple rows of data in tabular form",
	Long: 	 `use list to list everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		mseven, err := cmd.Flags().GetBool("Mseven")
		if err != nil {
			fmt.Println("getBool err: ", err)
			return
		}
		if mseven {
			showSevenMany(args)
		} else {
			showMany(args)
		}
	},
}

func init() {
	RootCmd.AddCommand(cmdShow, cmdConn, cmdList)
	cmdConn.Flags().BoolP("Mseven", "m", false, "If ucpe belongs to 6.X platform, Please use it")
	cmdShow.Flags().BoolP("Mseven", "m", false, "If ucpe belongs to 6.X platform, Please use it")
	cmdList.Flags().BoolP("Mseven", "m", false, "If ucpe belongs to 6.X platform, Please use it")
}

func show(sn string) {
	net.Search(sn)
}

func showSeven(sn string) {
	net.SearchSeven(sn)
}

func showMany(sns []string) {
	net.SearchMany(sns)
}

func showSevenMany(sns []string) {
	net.SearchSevenMany(sns)
}

// Run cmd
func Run() {
	RootCmd.Execute()
}