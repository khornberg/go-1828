package main

import "fmt"

const (
	Reset       = "\033[0m"
	Cyan        = "\033[36m"
	YellowBold  = "\033[33;1m"
	Magenta     = "\033[35m"
	White       = "\033[37m"
)

var colors = map[string]string{
	"cyan": Cyan,
	"magenta": Magenta,
	"white": White,
}

func PrintWord(s string, color string) {
	outputStr := colors[color] + "%s" + Reset + "\n"
	fmt.Printf(outputStr, s)
}

func PrintHeader(s string) {
	outputStr := YellowBold + "%s" + Reset + "\n"
	fmt.Printf(outputStr, s)
}
