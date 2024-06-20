package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10)) //[5 6 7 8 9]
	fmt.Println(piscine.AppendRange(10, 5)) //[]
	fmt.Println(piscine.AppendRange(0, 9))
	fmt.Println(piscine.AppendRange(-9, -3))
	fmt.Println(piscine.AppendRange(-42, -3))
}
