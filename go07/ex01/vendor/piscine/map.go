package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, v := range a {
		if f(v) == true {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i <= n / i; i++ {
		if n % i == 0 {
			return false 
		}
	}
	return true
}
