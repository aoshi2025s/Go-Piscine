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

	link = listPushBack(link, 9)
	link = listPushBack(link, 6)
	link = listPushBack(link, 13)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)
	link = listPushBack(link, 3)
	link = listPushBack(link, 8)
	link = listPushBack(link, 1)
	link = listPushBack(link, 2)
	link = listPushBack(link, -5)

	PrintList(piscine.ListSort(link))
}
