package piscine

func AppendRange(min, max int) []int {
	nums := []int{}
	if min >= max {
		return nums
	}
	for i := min; i < max; i++ {
		nums = append(nums, i)
	}
	return nums
}
