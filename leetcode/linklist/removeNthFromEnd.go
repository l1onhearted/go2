// 删除链表的倒数第 N 个结点   https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dhead := &ListNode{0, head}
	f := dhead
	s := dhead
	for i := 0; i < n+1; i++ {
		f = f.Next
	}
	for f != nil {
		s = s.Next
		f = f.Next
	}
	s.Next = s.Next.Next
	return dhead.Next
}