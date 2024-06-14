package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := []int{5,4,3,2,1,0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)

	s2 := []int{3,4,9,1,2,0,5,3,2,5,8,9}
	piscine.SortIntegerTable(s2)
	fmt.Println(s2)
}
