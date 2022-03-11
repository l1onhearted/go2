//150. 逆波兰表达式求值

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, i := range tokens {
		t, err := strconv.Atoi(i)
		if err == nil {
			stack = append(stack, t)
		} else {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			c := 0
			switch i {
			case "+":
				c = a + b
			case "-":
				c = b - a
			case "*":
				c = a * b
			case "/":
				c = b / a
			}
			stack = append(stack, c)
		}
	}
	return stack[0]
}

func main() {
	t := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(t))
}
