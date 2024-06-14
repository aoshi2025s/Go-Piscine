package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IterativeFactorial(4))
	fmt.Println(piscine.IterativeFactorial(5))
	fmt.Println(piscine.IterativeFactorial(6))
	fmt.Println(piscine.IterativeFactorial(7))
	fmt.Println(piscine.IterativeFactorial(8))
	fmt.Println(piscine.IterativeFactorial(9))
	fmt.Println(piscine.IterativeFactorial(10))
	fmt.Println(piscine.IterativeFactorial(+10))
	fmt.Println(piscine.IterativeFactorial(-1))
	fmt.Println(piscine.IterativeFactorial(-13))
	fmt.Println(piscine.IterativeFactorial(100))
}
