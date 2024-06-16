package main

import (
	"fmt"
	"piscine"
)

// ANSIエスケープシーケンスで色を定義
const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {
	result1 := piscine.Compare("Hello!", "Hello!")
	result2 := piscine.Compare("Salut!", "lut!")
	result3 := piscine.Compare("Ola!", "Ol!")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	if result1 == 0 {
		fmt.Println(Green + "test1 passed" + Reset)
	} else {
		fmt.Println(Red + "test1 failed" + Reset)
	}
	if result2 == -1 {
		fmt.Println(Green + "test2 passed" + Reset)
	} else {
		fmt.Println(Red + "test2 failed" + Reset)
	}
	if result3 == 1 {
		fmt.Println(Green + "test3 passed" + Reset)
	} else {
		fmt.Println(Red + "test3 failed" + Reset)
	}
}
