package main

import (
	"fmt"
	"piscine"
)

func main() {
	
	link := &piscine.List{}

	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, "there")
	piscine.ListPushBack(link, "how are you")
	piscine.ListPushBack(link, 1)

	fmt.Println(piscine.ListAt(link.Head, 3).Data)
	fmt.Println(piscine.ListAt(link.Head, 1).Data)
	fmt.Println(piscine.ListAt(link.Head, 7))
	// fmt.Println(piscine.ListAt(link.Head, 7).Data) -> segv
}
