package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	in_word := false
	for i := range runes {
		if ((runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= '0' && runes[i] <= '9')) {
			if !in_word {
				if runes[i] >= 'a' && runes[i] <= 'z' {
					runes[i] -= 'a' - 'A'
				}
				in_word = true
			} else {
				if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] += 'a' - 'A'
				}
			}
		} else {
			in_word = false
		}
	}
	return string(runes)
}
