package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n <= 0 || n > rLen(runes) {
		return 0
	}
	return runes[n-1]
}

func rLen(r []rune) int {
	c := 0
	for range r {
		c++
	}
	return c
}
