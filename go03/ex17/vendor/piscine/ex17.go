package piscine

func AtoiBase(s string, base string) int {
	if !baseIsValid(base) {
		return 0
	}
	result := 0
	for _, r := range s {
		value := sToNumInBase(r, base)
		if value < 0 {
			return result
		}
		result = result * strLen(base) + value
	}
	return result
}

func sToNumInBase(r rune, base string) int {
	index := 0
	for ; index < strLen(base); index++ {
		if r == rune(base[index]) {
			return index
		}
	}
	return -1
}

func strLen(str string) int {
	c := 0
	for range str {
		c++
	}
	return c
}

func baseIsValid(base string) bool {
	if strLen(base) < 2 {
		return false
	}
	var existAscii [128]bool
	for i := 0; i < strLen(base); i++ {
		c := int(base[i])
		if existAscii[c] == true || base[i] == '+' || base[i] == '-' {
			return false
		}
		existAscii[c] = true
	}
	return true
}
