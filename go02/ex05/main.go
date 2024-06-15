package main

import(
	"fmt"
	"piscine"
)

func main() {
	fmt.Println("1:",piscine.Sqrt(1))
	fmt.Println("4:",piscine.Sqrt(4))
	fmt.Println("3:",piscine.Sqrt(3))
	fmt.Println("-1:",piscine.Sqrt(-1))
	fmt.Println("6:",piscine.Sqrt(6))
	fmt.Println("144:",piscine.Sqrt(144))
	fmt.Println("256:",piscine.Sqrt(256))
	fmt.Println("0:",piscine.Sqrt(0))
	fmt.Println("-42",piscine.Sqrt(-42))
	fmt.Println("225",piscine.Sqrt(225))
}
