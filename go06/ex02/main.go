package main

import (
	"os"
	"piscine"
)

func main() {
	params := os.Arg
	if piscine.CountArg(params) != 1 {
		pisicne.PrintStr("Too many arguments\n")
		return
	}
	piscine.DisplayFile(params[0])
}

