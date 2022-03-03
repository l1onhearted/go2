//链表相交    https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	acur, bcur := headA, headB
	alen, blen := 0, 0
	for acur != nil { //求链表长度，不使用next
		acur = acur.Next
		alen++
	}
	for bcur != nil {
		bcur = bcur.Next
		blen++
	}
	acur, bcur = headA, headB
	if alen > blen {
		for i := 0; i < (alen - blen); i++ {
			acur = acur.Next
		}
		for acur != nil {
			if acur == bcur {
				return acur.Next
			} else {
				acur = acur.Next
				bcur = bcur.Next
			}
		}
		return nil
	} else {
		for i := 0; i < (blen - alen); i++ {
			bcur = bcur.Next
		}
		for bcur != nil {
			if acur == bcur {
				return bcur
			} else {
				acur = acur.Next
				bcur = bcur.Next
			}
		}
		return nil
	}
}