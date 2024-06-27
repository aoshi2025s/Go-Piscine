package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) == true {
			return true 
		}
	}
	return false
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
