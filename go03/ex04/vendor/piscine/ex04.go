package piscine

func Compare(a, b string) int {

	for i := range a {
		if i >= strLen(b) {
			return 1
		}
		if a[i] != b[i] {
			if a[i] < b[i] {
				return -1
			}
			return 1
		}
	}
	if strLen(a) < strLen(b) {
		return -1
	}
	return 0
}

func strLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}

