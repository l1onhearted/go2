//110.平衡二叉树
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isBalanced(root *TreeNode) bool {
// 	return height(root) >= 0
// }

// func height(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	lefth := height(root.Left)
// 	righth := height(root.Right)
// 	if lefth == -1 || righth == -1 || abs(lefth-righth) > 1 {
// 		return -1
// 	}
// 	return max(lefth, righth) + 1
// }
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	lefth := maxdepth(root.Left)
	righth := maxdepth(root.Right)
	if abs(lefth-righth) > 1 { //左右子树高度差大于1
		return false
	}
	return true
}

func maxdepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxdepth(root.Left), maxdepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
