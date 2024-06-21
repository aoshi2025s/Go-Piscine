package main

import (
	"os"
	"piscine"
)

func main() {
	params := os.Args
	if piscine.CountArg(params) > 2 {
		piscine.PrintStr("Too many arguments\n")
		return
	} else if piscine.CountArg(params) < 2 {
		piscine.PrintStr("File name missing\n")
		return
	}
	piscine.DisplayFile(params[0])
}

