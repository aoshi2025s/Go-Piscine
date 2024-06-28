package piscine

import (
	"os"
	"ft"
)

func CatWithStdin() {
	var buf [1]byte
	for {
		n, err := os.Stdin.Read(buf[:])
		if err != nil || n == 0 {
			break
		}
		ft.PrintRune(rune(buf[0]))
	}
}

func CatWithFile(fileName string) {
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

func CountArgs(args []string) int {
	c := 0
	for range args {
		c++
	}
	return c
}

/*
func CatWithFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		PrintStr("ERROR: " + err.Error() + "\n")
		return
	}

	defer file.Close()

	var buf [1]byte
	for {
		n, err := file.Read(buf[:])
		if err != nil || n == 0 {
			break
		}
		ft.PrintRune(rune(buf[0]))
	}
}
*/
