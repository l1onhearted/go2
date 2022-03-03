//349.两个数组的交集

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int { //暴力解法
	table := make(map[int]int)
	sec := []int{}
	for i, _ := range nums1 {
		if table[nums1[i]] == 0 {
			table[nums1[i]]++
		}
	}
	for i, _ := range nums2 {
		if table[nums2[i]] == 0 {
			continue
		}
		table[nums2[i]]--
	}
	for i := range table {
		if table[i] <= 0 {
			sec = append(sec, i)
		}
	}
	fmt.Println(table)
	return sec
}

func intersection1(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	res := []int{}
	for _, v := range nums2 {
		if c, ok := m[v]; c > 0 && ok {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

func main() {
	var nums1 = []int{1, 2, 2, 1}
	var nums2 = []int{2}
	fmt.Println(intersection(nums1, nums2))
}
