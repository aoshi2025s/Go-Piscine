package piscine

import (
	"ft"
	"os"
)

func DisplayFile(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		PrintStr("ERROR: " + err.Error() + "\n")
		return
	}
	PrintStr(string(content))
}

func PrintStr(str string) {
	for _, r := range str {
		ft.PrintRune(r)
	}
}

func CountArg(params []string) int {
	c := 0
	for range params {
		c++
	}
	return c
}
