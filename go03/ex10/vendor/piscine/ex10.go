package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	for r := range runes {
		if runes[r] == ' ' {
			continue
		} else if runes[r] >= 'A' && runes[r] <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
