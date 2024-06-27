package main

import (
	"fmt"
	"piscine"
)

func main() {
	
	link := &piscine.List{}

	piscine.List(link, "Hello")
	piscine.List(link, "there")
	piscine.List(link, "how are you")
}
