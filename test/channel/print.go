package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(chan int)
	go gr1(m)
	go gr2(m)
	time.Sleep(1e9)
}

func gr1(p chan int) {
	for i := 1; i <= 100; i++ {
		<-p
		if i%2 == 1 {
			fmt.Println("gr1:", i)
			time.Sleep(1e7)
		}
	}
}

func gr2(p chan int) {
	for i := 1; i <= 100; i++ {
		p <- 1
		if i%2 == 0 {
			fmt.Println("gr2:", i)
			time.Sleep(1e7)
		}
	}
}
