package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	// a<-1
	// go func() {
	// 	a <- 1
	// }()
	// go prA()
	go func() {
		for i := 0; i < 10; i++ {
			// select {
			// case <-a:
			// 	fmt.Println("A")
			// 	b <- 1
			// case <-m:
			// 	fmt.Println("A")
			// 	m<-1
			<-a
			fmt.Println("A")
			b <- 1
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-b
			fmt.Println("B")
			c <- 1
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-c
			fmt.Println("C")
			a <- 1
		}
		wg.Done()
	}()
	// go prB()
	// go prC()
	wg.Wait()
}

// func prA(){

// 	for i:=0;i<10;i++{
// 		<-m
// 		fmt.Println("A")
// 	}
// 	wg.Done()
// }
// func prB(){

// 	for i:=0;i<10;i++{
// 		<-m
// 		fmt.Println("B")
// 	}
// 	wg.Done()
// }
// func prC(){

// 	for i:=0;i<10;i++{
// 		<-m
// 		fmt.Println("C")
// 	}
// 	wg.Done()
// }
