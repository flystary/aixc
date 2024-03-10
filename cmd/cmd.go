package cmd

import (
	color "aixc/utils/color"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

}

var rootCmd = &cobra.Command{
	Use: "aixc",
}

var (
	Va  = color.Cyan("valor")
	Ta  = color.Cyan("tassadar")
	Wa  = color.Cyan("watsons")
	Wha = color.Cyan("watsonsha")

	Mo           = color.Cyan("model")
	Ve           = color.Cyan("version")
	Po           = color.Cyan("pop")
	En           = color.Cyan("enterprise")
	Pr           = color.Cyan("port")
	Up           = color.Cyan("update")
	selectOption = fmt.Sprintf("%s|%s|%s|%s|%s|%s", Mo, Ve, Po, En, Pr, Up)
	modeOption   = fmt.Sprintf("%s|%s|%s|%s", Va, Ta, Wa, Wha)
)

// Run cmd
func Run() {
	rootCmd.Execute()
}
