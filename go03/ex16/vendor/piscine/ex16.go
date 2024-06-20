package piscine

import "ft"

func PrintNbrBase(nbr int, base string) {
	if !baseIsValid(base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	baseLen := strLen(base)
	if nbr < 0 {
		ft.PrintRune('-')
		nbr = -nbr
	}
	if nbr == 0 {
		ft.PrintRune('0')
	}
	var result string
	for nbr > 0 {
		result = string(base[nbr % baseLen]) + result
		nbr /= baseLen
	}
	for _, r := range []rune(result) {
		ft.PrintRune(r)
	}
}

func strLen(str string) int {
	c := 0
	for range str {
		c++
	}
	return c
}

func baseIsValid(base string) bool {
	if strLen(base) < 2 {
		return false
	}
	var existAscii [128]bool
	for i := 0; i < strLen(base); i++ {
		c := int(base[i])
		if existAscii[c] == true || base[i] == '+' || base[i] == '-' {
			return false
		}
		existAscii[c] = true
	}
	return true
}
