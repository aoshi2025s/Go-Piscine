package piscine

func Median(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	nums = sortNums(nums)
	return nums[2]
}

func sortNums(nums []int) []int {
	for i := 0; i < elemLen(nums) - 1; i++ {
		for j := i + 1; j < elemLen(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}


func elemLen(n []int) int {
	c := 0
	for range n {
		c++
	}
	return c
}
