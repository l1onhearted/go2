//构建链表练习

package main

type MyLinkedList struct {
	val  int
	next *MyLinkedList
	pre  *MyLinkedList
}

func Constructor() MyLinkedList {
	dHead := new(MyLinkedList)
	dHead.next = dHead
	dHead.pre = dHead
	return *dHead //返回一个虚拟头节点
}

func (this *MyLinkedList) Get(index int) int {
	head := this.next
	for head != this && index > 0 {
		index--
		head = head.next
	}
	if index != 0 {
		return -1
	} else {
		return head.val
	}
}

func (this *MyLinkedList) AddAtHead(val int) {
	nhead := &MyLinkedList{val, this.next, this}
	this.next.pre = nhead
	this.next = nhead
}

func (this *MyLinkedList) AddAtTail(val int) {
	nhead := &MyLinkedList{val, this, this.pre}
	this.pre.next = nhead
	this.pre = nhead
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	head := this.next
	for head != this && index > 0 {
		index--
		head = head.next
	}
	if index > 0 {
		return
	} else {
		node := &MyLinkedList{val, head, head.pre}
		head.pre.next = node
		head.pre = node
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	head := this.next
	for head != this && index > 0 {
		index--
		head = head.next
	}
	if index > 0 {
		return
	} else {
		head.next.pre = head.pre
		head.pre.next = head.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
