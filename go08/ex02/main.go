package main

import (
	"os"
	"piscine"
)

func main() {
	args := os.Args
	if piscine.ElemLen(args) < 2 {
		return
	}
	params := args[1:]

	piscine.ComCheck(params)
}
