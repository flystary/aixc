package cmd

import (
	"github.com/spf13/cobra"

	c "aixc/control"
)

func init() {}

var rootCmd = &cobra.Command{
	Use: "aixc",
}

func show(sn string) {
	c.Search(sn)
}

func showSeven(sn string) {
	// c.SearchSeven(sn)
}

func showMany(sns []string) {
	c.SearchMany(sns)
}

func showSevenMany(sns []string) {
	// c.SearchSevenMany(sns)
}

func showManybyMode(mode string){
	println("此命令暂不支持!")
}
// Run cmd
func Run() {
	rootCmd.Execute()
}