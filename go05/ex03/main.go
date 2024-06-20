package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(""))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("norminette is the best cat "))
}

// go mod init ex03
// go run .
// []string{"Hello", "how", "are", "you?"}
