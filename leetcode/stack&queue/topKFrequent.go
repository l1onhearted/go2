//347. 前 K 个高频元素

package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	map1 := make(map[int]int, 0)
	for _, i := range nums {
		map1[i] += 1
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range map1 {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, 2)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return res
}

type IHeap [][2]int //构建小顶堆
func (h IHeap) Len() int {
	return len(h)
}
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h IHeap) Swap(i, j int) { //交换
	h[i], h[j] = h[j], h[i]
}
func (h *IHeap) Push(x interface{}) {
	// *h = append(*h, x.([2]int)) //空接口转换
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1] //先将x保存再改变堆大小
	*h = old[:n-1]
	return x
}
