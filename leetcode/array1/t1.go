package main

import "fmt"

func main() {
	var a = [][]int{{1, 2, 3}, {2, 4, 5}, {4, 5, 7}}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Print(a[i][j],",")
			fmt.Println(&a[i][j])
		}
	}
}
