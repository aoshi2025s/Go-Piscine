package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n <= 0 || n > strLen(s) {
		return 0
	}
	return runes[n-1]
}

func strLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}
