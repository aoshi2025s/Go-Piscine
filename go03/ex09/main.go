package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("012345"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
	fmt.Println(piscine.IsNumeric("01 0203"))
	fmt.Println(piscine.IsNumeric(""))
}
