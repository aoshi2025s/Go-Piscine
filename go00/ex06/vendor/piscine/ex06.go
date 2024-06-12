package piscine

import "ft"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var combs [10]int
	printCombR(combs[:n], n, 0)
	ft.PrintRune('\n')
}

func printCombR(combs []int, n int, index int) {
	if index == n {
		for i := 0; i < n; i++ {
			ft.PrintRune(rune(combs[i]) + '0')
		}
		if combs[0] != 10-n {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		return
	}
	start := 0
	if index > 0 {
		start = combs[index-1] + 1
	}
	for i := start; i <= 10-(n-index); i++ {
		combs[index] = i
		printCombR(combs, n, index+1)
	}
}
