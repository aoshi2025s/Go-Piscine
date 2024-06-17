package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			continue
		} else if runes[i] >= 'a' && runes[i] <= 'z' {
			continue
		} else if runes[i] >= '0' && runes[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
