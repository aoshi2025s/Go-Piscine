package piscine

func Sqrt(n int) int {
	i := 1
	for ; i < n / i; {
		i++;
	}
	if i * i == n {
		return i
	}
	return 0
}
