package main

import (
	"piscine"
	"ft"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, a)
	ft.PrintRune('\n')
	b := []int{9223372036854775807, -9223372036854775808, 0}
	piscine.ForEach(piscine.PrintNbr, b)
	ft.PrintRune('\n')
}
