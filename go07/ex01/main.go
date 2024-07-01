package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, a)
	fmt.Println(a)
	fmt.Println(result)

	b := []int{57, 13, 17, 19, 23, 2147483647}
	result = piscine.Map(piscine.IsPrime, b)
	fmt.Println(b)
	fmt.Println(result)

	c := []int{-9223372036854775808, 9223372036854775807}
	result = piscine.Map(piscine.IsPrime, c)
	fmt.Println(c)
	fmt.Println(result)
}
