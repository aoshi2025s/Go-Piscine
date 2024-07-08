package main

import (
	"fmt"
	"piscine"
)

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func main() {
	
	link := &piscine.List{}
	link2 := &piscine.List{}

	piscine.ListPushBack(link, "a")
	piscine.ListPushBack(link, "b")
	piscine.ListPushBack(link, "c")
	piscine.ListPushBack(link, "d")
	fmt.Println("---------first list----------")
	PrintList(link)

	piscine.ListPushBack(link2, "e")
	piscine.ListPushBack(link2, "f")
	piscine.ListPushBack(link2, "g")
	piscine.ListPushBack(link2, "h")
	fmt.Println("---------second list---------")
	PrintList(link2)

	fmt.Println("---------merged list---------")
	piscine.ListMerge(link, link2)
	PrintList(link)
}
