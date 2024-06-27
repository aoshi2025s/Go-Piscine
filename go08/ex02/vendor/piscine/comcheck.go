package piscine

import "ft"

func ComCheck(params []string) {
	for _, param := range params {
		if param == "42" || param == "piscine" || param == "piscine 42" {
			PrintAlert()
			return
		}
	}
}

func PrintAlert() {
	alertMsg := "Alert!!!"
	for _, r := range alertMsg {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func ElemLen(params []string) int {
	c := 0
	for range params {
		c++
	}
	return c
}
