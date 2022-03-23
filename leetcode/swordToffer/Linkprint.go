//06.从尾到头打印链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	queue := []int{}
	for head.Next != nil {
		queue = append(queue[], head.Val)
		head = head.Next
	}
	return queue

}
