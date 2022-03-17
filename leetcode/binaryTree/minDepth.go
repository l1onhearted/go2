package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int { //递归
	var getD func(node *TreeNode) int
	getD = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ld := getD(node.Left)
		rd := getD(node.Right)
		if node.Left == nil && node.Right != nil {
			return 1 + rd
		}
		if node.Left != nil && node.Right == nil {
			return 1 + ld
		}
		var min int
		if ld < rd {
			min = ld
		} else {
			min = rd
		}
		return 1 + min
	}
	return getD(root)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	min := 1
	for len(queue) > 0 {
		len := len(queue)
		for i := 0; i < len; i++ {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return min
			}
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		min += 1
	}
	return min
}
