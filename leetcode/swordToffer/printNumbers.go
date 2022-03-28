package main

import "math"

func printNumbers(n int) []int {
	r := []int{}
	for i := 1; i < int(math.Pow(10, float64(n))); i++ {
		r = append(r, i)
	}
	return r
}
