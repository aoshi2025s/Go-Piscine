package piscine

func ToUpper(str string) string {
	s := []rune(str)
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			s[i] -= ('a' - 'A')
		}
	}
	return string(s)
}
