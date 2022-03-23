package main

//11.旋转数组最小数字
func minArray(numbers []int) int {
	l, h := 0, len(numbers)-1
	for l < h {
		p := (l + h) / 2
		if numbers[p] < numbers[h] {
			h = p //考虑只有两个元素情况
		} else if numbers[p] > numbers[h] {
			l = p + 1
		} else {
			h = h - 1
		}
	}
	return numbers[h]
}
