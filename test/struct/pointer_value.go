package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B // b1是值
	// b1.thing = 2
	fmt.Println(b1)
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	fmt.Println(b2)
	b2.change()
	fmt.Println(b2.write())
}

/* 输出：
{1}
{1}
*/
