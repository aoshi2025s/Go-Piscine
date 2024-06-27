package piscine

func Rot14(s string) string {
	result := ""
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result += string(lower[(r - 'a' + 14) % 26])
		} else if r >= 'A' && r <= 'Z' {
			result += string(upper[(r - 'A' + 14) % 26])
		} else {
			result += string(r)
		}
	}
	return result
}
