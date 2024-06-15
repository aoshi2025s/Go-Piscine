package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println("1", piscine.IsPrime(1))
	fmt.Println("2", piscine.IsPrime(2))
	fmt.Println("3", piscine.IsPrime(3))
	fmt.Println("4", piscine.IsPrime(4))
	fmt.Println("21", piscine.IsPrime(21))
	fmt.Println("23", piscine.IsPrime(23))
	fmt.Println("57", piscine.IsPrime(57))
	fmt.Println("2147483647", piscine.IsPrime(2147483647))
	fmt.Println("0", piscine.IsPrime(0))
	fmt.Println("-42", piscine.IsPrime(-42))
	fmt.Println("2147483648", piscine.IsPrime(2147483648))
	fmt.Println("5", piscine.IsPrime(5))
	fmt.Println("13", piscine.IsPrime(13))
	fmt.Println("563", piscine.IsPrime(563))
}
