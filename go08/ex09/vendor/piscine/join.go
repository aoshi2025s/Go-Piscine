package piscine

func Join(strs []string, sep string) string {
	var result string
	for i := 0; i < elemLen(strs); i++ {
		if i != elemLen(strs) - 1 {
			result += strs[i]
			result += sep
		} else {
			result += strs[i]
		}
	}
	return result
}

func elemLen(strs []string) int {
	c := 0
	for range strs {
		c++
	}
	return c
}
