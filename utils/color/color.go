package color

import "github.com/fatih/color"


func Red(iput string) string {
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	return red(iput)
}

func Green(iput string) string {
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	return green(iput)
}

func Cyan(iput string) string {
	cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
	return cyan(iput)
}

func Blue(iput string) string {
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	return blue(iput)
}

func White(iput string) string {
	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	return white(iput)
}
