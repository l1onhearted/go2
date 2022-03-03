package main

import (
	"fmt"
	"time"
)

func main() {
	var nums = []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArraylen(7, nums))
}

func minSubArraylen(target int, nums []int) int {
	l, r, sublength := 0, 0, 0 //滑动窗口起始位置，终止位置，窗口长度
	sum := 0                   //滑动窗口数值之和
	t := time.Now().Unix()
	result := int(t)
	for r = 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target { //滑动窗口数值大则窗口右边不动左边缩小
			sublength = (r - l + 1)
			if result > sublength {
				result = sublength
			}
			sum -= nums[l] //窗口右移的同时减去移除元素
			l++
		}
		sublength = (r - l + 1)
	}
	if result == int(t) {
		return -1
	}
	return result
}
