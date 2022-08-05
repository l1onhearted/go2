package main

 type ListNode struct {
     Val int
     Next *ListNode
 }
//链表插入排序
 func insertionSortList(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	dHead:=&ListNode{Next:head}
	lsort:=head
	cur:=head.Next
	for cur!=nil{
		if lsort.Val<=cur.Val{
			lsort=lsort.Next
		}else{
			pre:=dHead
			for pre.Next.Val<=cur.Val{
				pre=pre.Next
			}
			lsort.Next=cur.Next
			cur.Next=pre.Next
			pre.Next=cur
		}
		cur=lsort.Next
	}
	return dHead.Next
}