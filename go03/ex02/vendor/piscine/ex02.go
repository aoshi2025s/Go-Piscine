package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	return runes[strLen(runes) - 1]
}

func strLen(runes []rune) int {
	c := 0
	for range runes {
		c++
	}
	return c 
}
