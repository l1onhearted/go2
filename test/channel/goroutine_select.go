package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	suck(ch1, ch2)

}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i*2 - 1
	}
}
func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("received ch1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("received ch2: %d\n", v)
		}
		time.Sleep(1e9)
	}
}
