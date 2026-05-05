package cmd

import (
	color "aixc/tools/color"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "aixc",
}

var (
	models       = []string{"model", "version", "pop", "enterprise", "port", "update"}
	modes        = []string{"valor", "tassadar", "watsons", "watsonsha"}
	selectOption = color.Cyan(strings.Join(models, "|"))
	modeOption   = color.Cyan(strings.Join(modes, "|"))

	// write *Write
	cli *CLI
)

func init() {

	cli = &CLI{
		mode:  "valor",
		seven: false,
		isOk:  false,
		Write: &Write{
			isWrite: false,
		},
		Options: &Options{
			Model:      &Model{isModel: false},
			Version:    &Version{isVersion: false},
			Pop:        &Pop{isPop: false},
			Enterprise: &Enterprise{isEnterprise: false},
			Port:       &Port{isPort: false},
			Update:     &Update{isUpdate: false},
		},
	}

}

// Run cmd
func Run() {
	rootCmd.Execute()

}
