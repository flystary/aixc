package color

import "github.com/fatih/color"

var (
	red   = color.New(color.FgRed, color.Bold)
	green = color.New(color.FgGreen, color.Bold)
	cyan  = color.New(color.FgCyan, color.Bold)
	blue  = color.New(color.FgBlue, color.Bold)
	white = color.New(color.FgWhite, color.Bold)
)

func Red(s string) string   { return red.Sprint(s) }
func Green(s string) string { return green.Sprint(s) }
func Cyan(s string) string  { return cyan.Sprint(s) }
func Blue(s string) string  { return blue.Sprint(s) }
func White(s string) string { return white.Sprint(s) }
