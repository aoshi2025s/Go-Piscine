package main

import  (
	"piscine"
	"fmt"
)

func main() {
	median := piscine.Median(2, 3, 8, 5, 7)

	fmt.Println(median)

	median = piscine.Median(8, 3, 7, 2, 5)
	fmt.Println(median)

	median = piscine.Median(1, 1, 5, 2, 3)
	fmt.Println(median)
}
