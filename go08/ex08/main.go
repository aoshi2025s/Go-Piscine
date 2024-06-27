package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)

	b := []int{343, 2342352, 25, 23423, 234523, 62352352352, 5235252}
	max = piscine.Max(b)
	fmt.Println(max)

	c := []int{}
	max = piscine.Max(c)
	fmt.Println(max)
}
