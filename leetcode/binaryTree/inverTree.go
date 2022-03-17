//226.翻转二叉树
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var tran func(node *TreeNode)
	tran = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		tran(node.Left)
		tran(node.Right)
	}
	tran(root)
	return root
}
