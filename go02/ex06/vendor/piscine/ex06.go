package piscine

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n / i; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}
