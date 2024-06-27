package piscine

func SortWordArr(a []string) {
	for i := 0; i < elemLen(a) - 1; i++ {
		for j := i + 1; j < elemLen(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func elemLen(strs []string) int {
	c := 0
	for range strs {
		c++
	}
	return c
}
