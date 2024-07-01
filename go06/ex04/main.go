package main

import (
	"ft"
	"os"
)

func main() {
	args := os.Args

	if piscine.Length(args) == 0 {
		piscine.PrintError("error")
		return
	}

	if args[0] == "-c" {
		if piscine.Length(args) < 3 {
			piscine.PrintError("error")
			return
		}

		count, err := piscine.Atoi(args[1])
		if !err {
			piscine.PrintError("Invalid number")
		}

		for i := 2; i < piscine.Length(args); i++ {
			if i > 2 {
				piscine.PrintSeparator()
			}
			piscine.TailFile(args[i], count)
		}
	} else {
		for i := 0; i < piscine.Length(args); i++ {
			if i > 0 {
				piscine.PrintSeparator()
			}
			piscine.TailFile(args[i], 10)
		}
	}
}
