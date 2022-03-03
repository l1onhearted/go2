package main

import "fmt"

func remove_e(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func main() {
	var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(remove_e(nums, 2))
	fmt.Println(nums)
}
