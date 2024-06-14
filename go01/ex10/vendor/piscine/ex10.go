package piscine

func Atoi(s string) int {
	result := 0
	index := 0
	length := strLen(s)
	sign := 1
	if (length > 0) && ((s[index] == '+') || (s[index] == '-')) {
		if (s[index] == '-') {
			sign = -1
		}
		index++;
	}
	for ; index < length; index++ {
		if (s[index] < '0' || s[index] > '9') {
			return (0)
		}
		result = result * 10 + int(s[index] - '0')
	}
	return (sign * result)
}

func strLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return (length)
}
