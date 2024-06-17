package main

import (
	"piscine"
	"ft"
)

func main() {
	piscine.PrintNbrBase(125, "0123456789") // 125
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01") // -1111101
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF") // 7D
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi") // -uoi
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa") // NV (????)
	ft.PrintRune('\n')
}
