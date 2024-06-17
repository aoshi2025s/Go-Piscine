package piscine

func IsNumeric(s string) bool {
	runes := []rune(s)
	for r := range runes {
		if runes[r] >= '0' && runes[r] <= '9' {
			continue
		} else if runes[r] == ' ' {
			continue
		} else {
			return false
		}
	}
	return true
}
