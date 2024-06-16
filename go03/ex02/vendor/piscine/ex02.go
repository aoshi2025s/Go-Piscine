package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	return runes[rLen(runes) - 1]
}

func rLen(r []rune) int {
	c := 0
	for range r {
		c++
	}
	return c 
}
