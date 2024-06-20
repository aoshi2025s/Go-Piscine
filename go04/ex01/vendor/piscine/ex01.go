package piscine

import "ft"

func PrintParams(params []string) {
	for i := 0; i < elemLen(params); i++ {
		for _, r := range params[i] {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}

func elemLen(strs []string) int {
	c := 0
	for range strs {
		c++
	}
	return c
}
