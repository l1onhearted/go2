package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//环形链表
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	low, fast := &ListNode{Next: head}, &ListNode{Next: head}
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		if fast == low {
			return true
		}
	}
	return false
}

//环形链表找出入幻节点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	f := false
	low, fast := &ListNode{Next: head}, &ListNode{Next: head}
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		if fast == low {
			f = true
			break
		}
	}
	if !f {
		return nil
	}
	cur := head
	for {
		if cur == low {
			return cur
		}
		cur = cur.Next
		low = low.Next
	}
}
