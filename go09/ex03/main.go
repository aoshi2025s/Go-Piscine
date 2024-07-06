package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}
	link2 := &piscine.List{}

	piscine.ListPushBack(link, "three")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "1")

	fmt.Println(piscine.ListLast(link)) // 1
	fmt.Println(piscine.ListLast(link2)) // <nil>
}
