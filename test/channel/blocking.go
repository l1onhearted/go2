package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int) //无缓冲通道发送就要接收否则死锁
	out <- 2
	go f1(out)
	// go func(out chan int) {
	// 	out <- 2
	// }(out)
	// go f1(out)
	// time.Sleep(1e9)
}
