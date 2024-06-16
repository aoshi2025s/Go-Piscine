package piscine

import "ft"

func PrintStr(s string) {
	runes := []rune(s)
	for _, c := range runes {
		ft.PrintRune(c)
	}
}
