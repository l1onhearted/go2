package main

func swapPairs(head *ListNode) *ListNode {
	dhead := new(ListNode) //虚拟头节点定义
	dhead.Next = head
	cur := dhead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp2 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp1
		cur.Next.Next.Next = tmp2
		cur = cur.Next.Next
	}
	return dhead.Next
}
