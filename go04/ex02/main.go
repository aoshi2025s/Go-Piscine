package main

import (
	"piscine"
	"os"
)

// go mod init ex01
// go run . norminette is the best cat

func main() {
	params := os.Args
	piscine.PrintReverseParams(params[1:])
}
