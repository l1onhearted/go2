package main

//46.全排列
func permute(nums []int) [][]int {
	var gen func(nums []int, index int, p []int, res [][]int, used []bool)
	gen = func(nums []int, index int, p []int, res [][]int, used []bool) {
		if index == len(nums) {
			temp := make([]int, len(p))
			copy(temp, p)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				p = append(p, nums[i])
				gen(nums, index+1, p, res, used)
				p = p[:len(p)-1]
				used[i] = false
			}
		}
		return
	}
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	gen(nums, 0, p, res, used)
	return res
}
