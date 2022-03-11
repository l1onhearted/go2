//前序迭代遍历 144
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			res = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		} // 向左遍历到最后,切换到右节点
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

