//20. 有效的括号

package main

func isValid(s string) bool {
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0) //数组模拟栈
	if s == " " {
		return true
	}
	for _, i := range []byte(s) {
		if i == '(' || i == '[' || i == '{' {
			stack = append(stack, i) //入栈
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
