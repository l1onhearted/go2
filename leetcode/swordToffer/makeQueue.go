package main

type CQueue struct {
	s1, s2 []int //定义一个输入栈,一个输出栈
}

func Constructor() CQueue {
	c := new(CQueue)
	return *c
}

func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.s1) == 0 && len(this.s2) == 0 {
		return -1
	} else if len(this.s2) == 0 { //输出栈为空则把输入栈push进输出栈
		for len(this.s1) > 0 { //
			this.s2 = append(this.s2, this.s1[0])
			this.s1 = this.s1[1:]
		}
	}
	t := this.s2[0]
	this.s2 = this.s2[1:]
	return t
}

// func main() {
// 	t := []int{1}
// 	fmt.Println(t[:1])
// }
