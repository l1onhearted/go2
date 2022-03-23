package main

//03.数组中重复数字
func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, i := range nums {
		if !m[i] {
			m[i] = true
			continue
		}
		if m[i] {
			return i
		}
	}
	return 0
}
