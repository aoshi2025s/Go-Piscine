package piscine

func ToLower(s string) string{
	r := []rune(s)
	for i := range r {
		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] += ('a' - 'A')
		}
	}
	return string(r)
}
