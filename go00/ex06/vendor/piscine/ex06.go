package piscine

import "syscall"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	fd := syscall.Stdout
	combs := make([]int, n)
	for i := 0; i < n; i++ {
		combs[i] = i
	}
	printCombRecursive(fd, combs, n, 0)
	syscall.Write(fd, []byte("\n"))
}

func printCombRecursive(fd int, combs []int, n int, index int) {
	if index == n {
		for i := 0; i < n; i++ {
			syscall.Write(fd, []byte{byte(combs[i] + '0')})
		}
		if combs[0] != 10-n {
			syscall.Write(fd, []byte(", "))
		}
		return
	}
	start := 0
	if index > 0 {
		start = combs[index-1] + 1
	}
	for i := start; i <= 9-(n-index); i++ {
		combs[index] = i
		printCombRecursive(fd, combs, n, index+1)
	}
}
