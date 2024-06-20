package piscine

import "ft"

func PrintString(programName string) {
	for _, r := range programName {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
