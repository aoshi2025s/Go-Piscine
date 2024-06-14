package piscine

func SortIntegerTable(table []int) {
	n := arrayLength(table)
	for i := 0; i < n; i++ {
		for j := 0; j < n - 1; j++ {
			if table[j] > table[j + 1] {
				table[j], table[j + 1] = table[j + 1], table[j]
			}
		}
	}
}

func arrayLength(array []int) int {
	length := 0
	for range array {
		length++
	}
	return (length)
}
