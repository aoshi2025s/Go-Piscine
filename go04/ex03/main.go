package main

import (
	"piscine"
	"os"
)

func main() {
	params := os.Args
	piscine.PrintSortedParams(params[1:])
}
