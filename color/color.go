package color

import "fmt"

var (
	red    = "\033[0;31m"
	green  = "\033[0;32m"
	reset  = "\033[0m"
	yellow = "\033[0;33m"
	blue   = "\033[0;34m"
)

func Red() string {
	return red
}

func Green() string {
	return green
}

func Yellow() string {
	return yellow
}

func Blue() string {
	return blue
}

func printColor(color string, text string) string {
	return color + text + reset
}

func Print(color string, text string) {
	fmt.Print(printColor(color, text))
}

func Println(color string, text string) {
	fmt.Println(printColor(color, text))
}
