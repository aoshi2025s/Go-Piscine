package piscine

func StrRev(s string) string {
	runes := []rune(s)
	length := strLen(s) - 1 
	for left, right := 0, length; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

func strLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
