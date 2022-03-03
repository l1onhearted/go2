package main

import "fmt"

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i, j := range nums {
		table[j] = i
	}
	for i, _ := range nums {
		t := target - nums[i]
		if i == table[t] {
			continue
		}
		if _, ok := table[t]; ok {
			r := []int{i, table[t]}
			return r
		}
	}
	return nil
}

func main() {
	nums1 := []int{3, 2, 4}
	fmt.Println(twoSum(nums1, 6))
}
