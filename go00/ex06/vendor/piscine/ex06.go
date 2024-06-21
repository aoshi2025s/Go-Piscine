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

/*
PrintCombN(3)
printCombR(combs[:3], 3, 0)
for i := 0; i <= 7; i++ {
	combs[0] = 0
	printCombR(combs, 3, 1)
	for i := 1; i <= 8; i++ {
		combs[1] = 1
		printCombR(combs, 3, 2)
		for i := 2; i <= 9; i++ {
			combs[2] = 2
			printCombR(combs, 3, 3)
			// 012,
			combs[2] = 3
			printCombR(combs, 3, 3)
			// 013,
			combs[2] = 4
			printCombR(combs, 3, 3)
			// 014,
			.
			.
			.
		}
	}
}

PrintCombN(1)
printCombR(combs[:1], 1, 0)
for i := 0; i <= 9; i++ {
	combs[0] = 0
	printCombR(combs, 1, 1)
	// 0, 
	combs[0] = 1
	printCombR(combs, 1, 1)
	// 1, 
	.
	.
	.
	combs[0] = 9
	printCombR(combs, 1, 1)
	// 9
}
*/
