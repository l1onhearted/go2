//239.滑动窗口最大值

package main

type myqueue struct {
	queue []int
}

func NewMyQueue() *myqueue {
	return &myqueue{
		queue: make([]int, 0),
	}
}
func (m *myqueue) Front() int { //返回队首
	return m.queue[0]
}
func (m *myqueue) Back() int { //返回队尾
	return m.queue[len(m.queue)-1]
}
func (m *myqueue) Empty() bool {
	return len(m.queue) == 0
}
func (m *myqueue) Push(val int) {
	for !m.Empty() && val > m.Back() { //队列不为空并且值大于队尾值
		m.queue = m.queue[:len(m.queue)-1] //删去队尾元素
	}
	m.queue = append(m.queue, val)
}
func (m *myqueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:] //删去队头元素
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front()) //记录前k个元素最大值
	for i := k; i < length; i++ {
		queue.Pop(nums[i-k])             //滑动窗口移除最前面的元素
		queue.Push(nums[i])              //滑动窗口添加最后面元素
		res = append(res, queue.Front()) // 记录最大值
	}
	return res
}
