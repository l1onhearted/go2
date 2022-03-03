// package main

package main

import "fmt"

func h_search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] == target {
			return middle
		} else if arr[middle] > target {
			right = middle - 1
		} else if arr[middle] < target {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	// x := min(1, 3, 2, 0)
	// fmt.Printf("The minimum is: %d\n", x)
	// slice := []int{7, 9, 3, 5, 1}
	// x = min(slice...)
	// fmt.Printf("The minimum in the slice is: %d", x)
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println(h_search(nums, 3))

}
