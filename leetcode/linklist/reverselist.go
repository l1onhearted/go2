//反转链表
package main

type ListNode struct {
	val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	pre := new(ListNode)
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return cur
}
