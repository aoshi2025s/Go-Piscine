package piscine

import (
	"ft"
)

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

func PrintNbr(n int) {
	if n < 0 {
		if n  == -(1 << 31) {
			PrintStr("-2147483648")
			return
		} else if n == -(1 << 63) {
			PrintStr("-9223372036854775808")
			return
		} else {
			ft.PrintRune('-')
			n = -n
		}
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	ft.PrintRune(rune(n % 10 + '0'))
}

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

/*
func PrintNbr(n int) {
	result := ""
	if n < 0 {
		ft.PrintRune('-')
		n = -n
	}
	for n >= 10 {
		result = string(n % 10 + '0') + result
		n /= 10
	}
	result = string(n % 10 + '0') + result
	for _, r := range result {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
*/


