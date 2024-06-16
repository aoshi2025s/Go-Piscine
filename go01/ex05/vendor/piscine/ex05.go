package piscine

func StrLen(s string) int {
	runes := []rune(s)
	c := 0
	for range runes {
		c++
	}
	return c	
}
