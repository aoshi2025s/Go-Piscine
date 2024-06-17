package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	for r := range runes {
		if runes[r] < 'A' || runes[r] > 'Z' {
			return false
		}
	}
	return true
}
