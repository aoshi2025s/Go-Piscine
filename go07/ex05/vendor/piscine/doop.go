package piscine

func CalcuDoop(v1, v2 int, op string) (int, bool) {
	if op[0] == '+' {
		if isSafeAdd(v1, v2) == false {
			return 0, false
		}
		return v1 + v2, true
	}
	if op[0] == '-' {
		if v1 < getMinInt() + v2 {
			return 0, false
		}
		return v1 - v2, true
	}
	if op[0] == '*' {
		if isSafeMultiply(v1, v2) == false {
			return 0, false
		}
		return v1 * v2, true
	}
	if op[0] == '/' {
		if isSafeDivide(v1, v2) == false {
			return 0, false
		}
		return v1 / v2, true
	}
	if op[0] == '%' {
		if v2 == 0 {
			return 0, false
		}
		return v1 % v2, true
	}
	return 0, false
}

func Atoi(s string) (int, bool) {
	result := 0
	minus := 1
	
	// https://github.com/aoshi2025s/Go-Piscine/issue/8
	// string と runeを比較しない
	if s[0] == '-' {
		minus = -1
		s = s[1:]
	}

	for _, r := range s {
		if IsOverFlow(result * minus, int(r - '0')) == true {
			return result * minus, false
		}
		result = result * 10 + int(r - '0')
	}
	return result * minus, true
}

func getMaxInt() int {
	if ^uint(0) >> 32 == 0 {
		// 32bit
		return 1 << 31 - 1 // 2147483647
	} else {
		// 64bit
		return 1 << 63 - 1 // 9223372036854775807
	}
}

func getMinInt() int {
	if ^uint(0) >> 32 == 0 {
		// 32 bit
		return -1 << 31 //-2147483648
	} else {
		// 64 bit
		return -1 << 63 // -9223372036854775808
	}
}


func IsOverFlow(n, p int) bool {
	maxInt := getMaxInt()
	minInt := getMinInt()
	
	// n * 10  + p > maxInt
	// n * 10 > maxInt - p
	// n > (maxInt - p) / 10
	
	// 下記だと現状ダメっぽい。
	/*
	if n > (maxInt - p) / 10 {
		return true
	}
	if n < (minInt - p) / 10 {
		return true
	}
	*/

	if n > maxInt / 10 {
		return true
	} else if n == maxInt / 10 && p > maxInt % 10 {
		return true
	}
	if n < minInt / 10 {
		return true
	} else if n == minInt / 10 && p >  -(minInt % 10) {
		return true
	}
	return false
}


func IsValidInput(params []string) bool {
	if elemLen(params) != 3 {
		return false
	}
	if isDigit(params[0]) == false || isDigit(params[2]) == false {
		return false
	}
	if isDoop(params[1]) == false {
		return false
	}
	return true
}

func isDoop(s string) bool {
	if strLen(s) > 1 {
		return false
	}
	runes := []rune(s)
	if runes[0] == '*' || runes[0] == '/' || runes[0] == '%' {
		return true
	} else if runes[0] == '+' || runes[0] == '-' {
		return true
	} else {
		return false
	}
}

func isSafeAdd(v1, v2 int) bool {
	maxInt := getMaxInt()
	minInt := getMinInt()

	if v1 > 0 && v2 > 0 && v2 > maxInt - v1 {
		return false
	}
	if v1 < 0 && v2 < 0 && v2 < minInt - v1 {
		return false
	}
	return true
}

func isSafeMultiply(v1, v2 int) bool {
	maxInt := getMaxInt()
	minInt := getMinInt()
	
	// INT_MIN * (-1)
	if (v1 == minInt && v2 == -1) || (v2 == minInt && v1 == -1) {
		return false
	}
	if v1 > 0 && v2 > 0 && v1 > maxInt / v2 {
		return false
	}
	if v1 > 0 && v2 < 0 && v2 < minInt / v1 {
		return false
	}
	if v1 < 0 && v2 > 0 && v1 < minInt / v2 {
		return false
	}
	if v1 < 0 && v2 < 0 && v1 < maxInt / v2 {
		return false
	}
	return true
}

func isSafeDivide(v1, v2 int) bool {
	minInt := getMinInt()

	if v2 == 0 {
		return false
	}

	if v1 == minInt && v2 == -1 {
		return false
	}
	return true
}

func strLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}

func isDigit(s string) bool {
	if s[0] == '-' {
		s = s[1:]
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func elemLen(strs []string) int {
	count := 0
	for range strs {
		count++
	}
	return count
}
