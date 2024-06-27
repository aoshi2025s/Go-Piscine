package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, str := range *ptr {
		if str != "" {
			count++
		}
	}
	result := make([]string, count)
	i := 0
	for _, str := range *ptr {
		if str != "" {
			result[i] = str
			i++
		}
	}
	*ptr = result
	return count
}
