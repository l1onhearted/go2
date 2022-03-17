//101.对称二叉树
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func isSymmetric(root *TreeNode) bool {
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		//空节点的情况
		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left == nil && right == nil {
			return true
		} else if left.Val != right.Val { //数值不相同的情况
			return false
		}
		outside := compare(left.Left, right.Right)
		inside := compare(left.Right, right.Left)
		return outside && inside
	}
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

//迭代使用队列
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		leftnode := queue[0]
		queue = queue[1:]
		rightnode := queue[0]
		queue = queue[1:]
		if leftnode == nil && rightnode == nil { //左右都为空则对称
			continue
		}
		if leftnode == nil || rightnode == nil || leftnode.Val != rightnode.Val {
			return false
		}
		queue = append(queue, leftnode.Left)
		queue = append(queue, rightnode.Right)
		queue = append(queue, leftnode.Right)
		queue = append(queue, rightnode.Left)
	}
	return true
}
