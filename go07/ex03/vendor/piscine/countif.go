package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, v := range tab {
		if f(v) == true {
			count++
		}
	}
	return count
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
