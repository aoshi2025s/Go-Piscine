package piscine

import (
	"ft"
)

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
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

func PrintNbr(n int) {
	if n < 0 {
		ft.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	ft.PrintRune(rune(n % 10 + '0'))
}
