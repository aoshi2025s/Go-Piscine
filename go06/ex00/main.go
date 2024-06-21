package main

import (
	"os"
	"piscine"
	"ft"
)

// even関数とEvenMsgとOddMsgとbooleanとyes,noのマクロ定義？が必要そう

func main() {
	params := os.Arg
	lengthOfArg := piscine.CountArg(params)
	if isEven(lengthOfArg) == {
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

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

// The necessary change must be applied for the program to work.
