package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		head = head.Next
	}
	cur := &ListNode{Next: head}
	for ; cur.Next != nil; cur = cur.Next {
		if cur.Next.Val == val {
			if cur.Next.Next == nil {
				cur.Next = nil
			} else if cur.Next == nil {
				cur = nil
			} else {
				cur.Next = cur.Next.Next
			}
		}
	}
	return head
}
