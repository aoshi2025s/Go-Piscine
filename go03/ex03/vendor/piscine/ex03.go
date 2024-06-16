package piscine

func Index(s string, toFind string) int {
	s_len := strLen(s)
	tf_len := strLen(toFind)

	for index := 0; index < s_len - tf_len + 1; index++ {
		exist := true
		for tf_index := 0; tf_index < tf_len; tf_index++ {
			if s[index + tf_index] != toFind[tf_index] {
				exist = false
			}
		}
		if exist {
			return index
		}
	}
	return -1
}

func strLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}
