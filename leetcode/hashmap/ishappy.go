//202.快乐数
package main

import "fmt"

func isHappy(n int) bool {

	table := make(map[int]int)
	for {
		if n == 1 {
			return true
		}
		sum := getsum(n)
		if table[sum] == 1 {
			return false
		} else {
			table[sum] = 1
		}
		n = sum
	}
}
func getsum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
func main() {
	fmt.Println(isHappy(19))
}
