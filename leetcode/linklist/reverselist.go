//反转链表
package main

type listNode struct {
	val  int
	next *ListNode
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
