package main

import (
	"piscine"
	"os"
)

// go run . 1 a 2 A 3 b 4 C
// 1 2 3 4 A C a b

func main() {
	params := os.Args
	piscine.PrintSortedParams(params[1:])
}
