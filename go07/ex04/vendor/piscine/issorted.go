package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if elemLen(a) < 2 {
		return true
	}

	for i := 0; i < elemLen(a) - 1; i++ {
		if f(a[i], a[i + 1]) == 1 {
			return false
		}
	}
	return true
}

func Compare(a, b int) int {
	if a <= b {
		return 0
	}
	return 1
}

func elemLen(a []int) int {
	count := 0
	for range a {
		count++
	}
	return count
}
