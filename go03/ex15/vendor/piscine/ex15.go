package piscine

func Join(strs []string, sep string) string {
	if strsLen(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, str := range strs[1:] {
		result += sep + str
	}
	return result
}

func strsLen(strs []string) int {
	c := 0
	for range strs {
		c++
	}
	return c
}
