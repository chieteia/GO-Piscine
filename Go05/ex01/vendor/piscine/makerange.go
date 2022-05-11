package piscine

func MakeRange(min, max int) []int {
	cap := max - min
	if cap <= 0 {
		return []int{}
	}
	nums := make([]int, cap)
	for i := 0; i < cap; i++ {
		nums[i] = min + i
	}
	return nums
}
