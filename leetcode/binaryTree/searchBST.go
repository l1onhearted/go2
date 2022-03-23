//700.二叉搜索树的搜索
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		if node.Val == val {
			return node
		} else if node.Val < val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return nil
}

//98验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var midp func(node *TreeNode)
	t := []int{}
	midp = func(node *TreeNode) {
		if node == nil {
			return
		}
		midp(node.Left)
		t = append(t, node.Val)
		midp(node.Right)
	}
	midp(root)
	for i := 0; i < len(t)-1; i++ {
		if t[i] >= t[i+1] {
			return false
		}
	}
	return true
}

//二叉搜索树最小绝对差
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var midp func(node *TreeNode)
	t := []int{}
	midp = func(node *TreeNode) {
		if node == nil {
			return
		}
		midp(node.Left)
		t = append(t, node.Val)
		midp(node.Right)
	}
	min := 100
	for i := 1; i < len(t); i++ {
		tmp := t[i] - t[i-1]
		if tmp < min {
			min = tmp
		}
	}
	return min
}

//701.二叉搜索树的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode { //递归
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
func insertIntoBST(root *TreeNode, val int) *TreeNode { //迭代
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	var pnode *TreeNode //记录前一个节点
	for node != nil {   //让node走到树的末端
		if val > node.Val {
			pnode = node
			node = node.Right
		} else {
			pnode = node
			node = node.Left
		}
	}
	if val > pnode.Val {
		pnode.Right = &TreeNode{Val: val}
	} else {
		pnode.Left = &TreeNode{Val: val}
	}
	return root
}
