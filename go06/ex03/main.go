package main

import (
	"piscine"
	"os"
)

func main() {
	args := os.Args[1:]

	if piscine.CountArgs(args) == 0 {
		piscine.CatWithStdin()
		return
	}
	for i := 0; i < piscine.CountArgs(args); i++ {
		piscine.CatWithFile(args[i])
	}
}
