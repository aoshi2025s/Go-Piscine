package piscine

func Split(s, sep string) []string {
	var result []string
	str := ""
	for i := 0; i < strLen(s); i++ {
		if i < strLen(s) - strLen(sep) && isSep(s[i:], sep) {
			i += strLen(sep) - 1
			result = append(result, str)
			str = ""
		} else {
			str += string(s[i])
		}
	}
	result = append(result, str)
	return result
}

func strLen(str string) int {
	c := 0
	for range str {
		c++
	}
	return c
}

func isSep(s, sep string) bool {
	for i := 0; i < strLen(sep); i++ {
		if sep[i] != s[i] {
			return false
		}
	}
	return true
}
