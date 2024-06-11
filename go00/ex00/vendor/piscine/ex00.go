package piscine

import "ft"

func PrintAlp() {
	for r := 'a'; r <= 'z'; r++ {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
