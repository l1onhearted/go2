package main

import "fmt"

func main() {
	var nums = []int{-4, -1, 0, 3, 10}
	fmt.Println(len(nums))
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	var nums1 []int = make([]int, len(nums))
	for i, j, k := 0, len(nums)-1, len(nums)-1; k > -1; { //双指针法,最大值分布在两侧
		l := nums[i] * nums[i]
		r := nums[j] * nums[j]
		if l > r {
			nums1[k] = l
			i++
			k--
		} else if l < r {
			nums1[k] = r
			j--
			k--
		} else {
			nums1[k] = r
			i++
			// j--      对称情况会报错   []
			k--
		}
	}
	return nums1
}
