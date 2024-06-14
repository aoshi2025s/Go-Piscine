package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IterativePower(2,3))
	fmt.Println(piscine.IterativePower(2,-3))
	fmt.Println(piscine.IterativePower(2,0))
	fmt.Println(piscine.IterativePower(2,31))
	fmt.Println(piscine.IterativePower(2,32))
	fmt.Println(piscine.IterativePower(2,63))
	fmt.Println(piscine.IterativePower(2,65))
	fmt.Println(piscine.IterativePower(15,2))
	fmt.Println(piscine.IterativePower(16,2))
	fmt.Println(piscine.IterativePower(99,99))
}
