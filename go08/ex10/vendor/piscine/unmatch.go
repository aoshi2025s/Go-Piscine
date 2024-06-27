package piscine

func Unmatch(a []int) int {
	for i := 0; i < elemLen(a); i++ {
		existPare := false
		for j := 0; j < elemLen(a); j++ {
			if i != j && a[i] == a[j] {
				existPare = true
			}
		}
		if existPare == false {
			return a[i]
		}
	}
	return -1
}

func elemLen(nums []int) int {
	c := 0
	for range nums {
		c++
	}
	return c
}
