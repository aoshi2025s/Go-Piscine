package piscine

func ConcatParams(strs []string) string {
	if elemLen(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, str := range strs[1:] {
		result += "\n" + str
	}
	return result
}

func elemLen(elems []string) int {
	c := 0
	for range elems {
		c++
	}
	return c
}
