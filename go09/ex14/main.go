package main

import (
	"fmt"
	"piscine"
)

func PrintList(l *piscine.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
	n := &piscine.NodeI{Data: data}

	if l == nil {
		return n
	}
	it := l
	for it.Next != nil {
		it = it.Next
	}
	it.Next = n
	return l
}

func main() {
	var link *piscine.NodeI
	var link2 *piscine.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)
	fmt.Println("-----link-------")
	PrintList(link)
	
	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)
	fmt.Println("------link2------")
	PrintList(link2)

	fmt.Println("------after merge-----")
	PrintList(piscine.SortedListMerge(link2, link))
}
