package main

import (
	"piscine"
	"os"
)

// go build . -> ./ex00
// go build -o main -> ./main

func main() {
	programName := os.Args[0]
	piscine.PrintString(programName)
}
