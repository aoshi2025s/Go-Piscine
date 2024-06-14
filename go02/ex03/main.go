package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursivePower(2,3))
	fmt.Println(piscine.RecursivePower(2,-3))
	fmt.Println(piscine.RecursivePower(2,0))
	fmt.Println(piscine.RecursivePower(2,31))
	fmt.Println(piscine.RecursivePower(2,32))
	fmt.Println(piscine.RecursivePower(2,63))
	fmt.Println(piscine.RecursivePower(2,65))
	fmt.Println(piscine.RecursivePower(15,2))
	fmt.Println(piscine.RecursivePower(16,2))
	fmt.Println(piscine.RecursivePower(99,99))
}
