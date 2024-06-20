package piscine

import "ft"

func PrintParams(params []string) {
	for _, param := range params {
		for _, r := range param {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}
