package main //交替输出偶数奇数

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	m := make(chan int)
	go gr1(m)
	go gr2(m)
	wg.Wait()
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func gr1(p chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		<-p
		if i%2 == 1 {
			fmt.Println("gr1:", i)
			time.Sleep(1e7)
		}
	}
}

func gr2(p chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		p <- 1
		if i%2 == 0 {
			fmt.Println("gr2:", i)
			time.Sleep(1e8)
		}
	}
}
