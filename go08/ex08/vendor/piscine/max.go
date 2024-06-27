package piscine

func Max(a []int) int {
	maxValue := 0
	for _, num := range a {
		if num > maxValue {
			maxValue = num
		}
	}
	return maxValue
}
