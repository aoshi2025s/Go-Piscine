package main

import (
	"os"
	"piscine"
	"ft"
)

type boolean int

const (
	yes boolean = 1
	no boolean = 0
)

const (
	EvenMsg = "I have an even number of arguments"
	OddMsg = "I have an odd number of arguments"
)

func main() {
	params := os.Args
	lengthOfArg := piscine.CountArg(params) - 1
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func even(nbr int) int {
	if nbr % 2 == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}
