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

	piscine.ListReverse(link)
	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
