package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println("count of active bits")
	fmt.Println(piscine.ActiveBits(7))
	fmt.Println(piscine.ActiveBits(4))
	fmt.Println(piscine.ActiveBits(133))
	fmt.Println(piscine.ActiveBits(991199))
	fmt.Println(piscine.ActiveBits(214746847))
	fmt.Println(piscine.ActiveBits(57))
	
	fmt.Println("bit of value")
	fmt.Printf("%b\n", 7)
	fmt.Printf("%b\n", 4)
	fmt.Printf("%b\n", 133)
	fmt.Printf("%b\n", 991199)
	fmt.Printf("%b\n", 214746847)
	fmt.Printf("%b\n", 57)
}
