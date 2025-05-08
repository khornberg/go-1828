package main

import "fmt"

const (
	Reset       = "\033[0m"
	Cyan        = "\033[36m"
)

func PrintWord(s string) {
	outputStr := Cyan + "%s" + Reset
	fmt.Printf(outputStr, s)
}

