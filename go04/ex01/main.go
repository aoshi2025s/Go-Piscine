package main

import (
	"piscine"
	"os"
)

func main() {
	params := os.Args
	piscine.PrintParams(params[1:])
}
