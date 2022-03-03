package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	if head != nil {
		dnode := new(ListNode)
		// t := new(ListNode)
		dnode.Next = head
		t := dnode
		for t.Next != nil {
			// t = dnode.Next
			if t.Next.Val == val {
				t.Next = t.Next.Next
			} else {
				t = t.Next
			}
		}
		return dnode.Next
	} else {
		return nil
	}
}
