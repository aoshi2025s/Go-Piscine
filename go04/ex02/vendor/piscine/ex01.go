package piscine

import "ft"

func PrintReverseParams(params []string){
	for i := elemLen(params) - 1; i >= 0; i-- {
		for _, r := range params[i] {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}

func elemLen(elems []string) int {
	c := 0
	for range elems {
		c++
	}
	return c
}
