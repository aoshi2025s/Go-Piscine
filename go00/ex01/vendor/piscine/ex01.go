package piscine

import "ft"

func PrintRevAlpha() {
	for r := 'z'; r >= 'a'; r-- {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
