//1047. 删除字符串中的所有相邻重复项
package main

import "fmt"

func removeDuplicates(s string) string {
	str := []byte(s)
	stack := make([]byte, 0)
	for _, i := range str {
		if len(stack) == 0 || stack[len(stack)-1] != i {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
func main() {
	str := "abbaca"
	fmt.Println(removeDuplicates((str)))
}
