package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursiveFactorial(4))
	fmt.Println(piscine.RecursiveFactorial(5))
	fmt.Println(piscine.RecursiveFactorial(6))
	fmt.Println(piscine.RecursiveFactorial(7))
	fmt.Println(piscine.RecursiveFactorial(8))
	fmt.Println(piscine.RecursiveFactorial(9))
	fmt.Println(piscine.RecursiveFactorial(10))
	fmt.Println(piscine.RecursiveFactorial(+10))
	fmt.Println(piscine.RecursiveFactorial(-1))
	fmt.Println(piscine.RecursiveFactorial(-13))
	fmt.Println(piscine.RecursiveFactorial(100))
}
