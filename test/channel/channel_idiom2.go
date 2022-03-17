package main //使用for range从通道中获取数据

import (
	"fmt"
	"time"
)

func main() {
	suck(pump())
	time.Sleep(1e7)
}

func pump() (chan int, chan int) {
	ch := make(chan int)
	ch1 := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i * 2
			ch1 <- i*2 - 1
		}
	}()
	return ch, ch1
}

func suck(ch, ch1 chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	go func() {
		for v := range ch1 {
			fmt.Println(v)
		}
	}()
}
