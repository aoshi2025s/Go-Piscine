package piscine

import "ft"

func PrintSortedParams(params []string) {
	for i := 0; i < elemLen(params) - 1; i++ {
		for j := 0; j < elemLen(params) - 1; j++ {
			if params[j] > params[j + 1] {
				params[j], params[j + 1] = params[j + 1], params[j]
			}
		}
	}

	for _, param := range params {
		for _, r := range param {
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
