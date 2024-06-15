package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(4))
	fmt.Println(piscine.FindNextPrime(2))
	fmt.Println(piscine.FindNextPrime(57))
	fmt.Println(piscine.FindNextPrime(-2))
	fmt.Println(piscine.FindNextPrime(-1))
	fmt.Println(piscine.FindNextPrime(0))
	fmt.Println(piscine.FindNextPrime(16))
	fmt.Println(piscine.FindNextPrime(46))
	fmt.Println(piscine.FindNextPrime(2147483646))
}
