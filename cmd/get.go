package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("type", "T", "","ucpe type, Please use it")
	getCmd.Flags().BoolP("entn", "E", false,"ucpe enterprise, Please use it")

}

var getCmd = &cobra.Command{
	Use:   	 "get",
	// Example: "xc conn 7x00114401917b5f0",
	Short:   "Print multiple rows of data in tabular form",
	Long: 	 `use list to list everything you want form cpe`,
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		var xtype string
		entn, err := cmd.Flags().GetBool("entn")
		if err != nil {
			println("getbool err: ", err)
			return
		}

		stype, err := cmd.Flags().GetString("type")
		if err != nil {
			println("getstring err: ", err)
			return
		}

		if stype == "" {
			xtype = "enterprise"
		}else{
			xtype = stype
		}
		println(entn)
		println(xtype)

		println(args[0])
	},
}