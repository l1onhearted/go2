package main

//倒数第k个节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	l, f := head, head
	for i := 0; i <= k; i++ {
		f = f.Next
	}
	for f.Next != nil {
		l = l.Next
		f = f.Next
	}
	return l
}
