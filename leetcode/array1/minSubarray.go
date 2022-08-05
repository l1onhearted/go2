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
	s := []int{1, 2, 3, 4}
	s = make([]int)
	s = make([]int, 0)
	s = make([]int, 5, 10)
}

// func add(a, b string) string {
// 	// 补充具体实现
// 	carry := 0
// 	i := len(a) - 1
// 	j := len(b) - 1
// 	var x, y int
// 	var res string
// 	for i >= 0 || j >= 0 || carry != 0 {
// 		if i >= 0 {
// 			x = getint(a[i])
// 		} else {
// 			x = 0
// 		}
// 		if j >= 0 {
// 			y = getint(b[j])
// 		} else {
// 			y = 0
// 		}
// 		temp := x + y + carry
// 		res += string(getChar(temp % 36))
// 		carry = temp / 36
// 		i--
// 		j--
// 	}
// 	res=rev(res)
// 	return res
// }

// func getint(ch byte) int {
// 	if '0' <= ch && ch <= '9' {
// 		return int(ch - '0')
// 	} else {
// 		return int(ch - 'a' + 10)
// 	}
// }

// func getChar(n int) byte {
// 	if n <= 9 {
// 		return byte(n) + '0'
// 	} else {
// 		return byte(n-10) + 'a'
// 	}
// }
// func rev(str string) string {
// 	resstr := []byte(str)
// 	for i, j := 0, len(str)-1; i < j; {
// 		resstr[i], resstr[j] = resstr[j], resstr[i]
// 	}
// 	return string(resstr)
// }

m1 := make(map[int]int)
for i, n := range nums {
	n1, f := m1[target-n]
	if f {
		return []int{n1, i}
	} else {
		m1[n] = i
	}
}
return nil