package main

import (
	"fmt"
	"piscine"
)

func main() {
	elem := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elem))
}
