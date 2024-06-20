package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	
	// baseFrom -> int
	decimal := AtoiBase(nbr, baseFrom)

	// int -> baseTo
	baseLen := strLen(baseTo)
	result := ""
	for decimal > 0 {
		result = string(baseTo[decimal % baseLen]) + result
		decimal /= baseLen
	}
	return result
}

func strLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}

func AtoiBase(s string, base string) int {
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
