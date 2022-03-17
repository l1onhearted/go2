//二叉树前序,中序,后序遍历 144.145.94

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	// var traversal func(node *TreeNode)
	func traversal(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	res = []int{}
	traversal(root)
	return res
}

func traversal(node *TreeNode) { //前序
	if node == nil {
		return
	}
	res = append(res, node.Val) //中
	traversal(node.Left)        //左
	traversal(node.Right)       //右
}

func traversal(node *TreeNode) { //中序
	if node == nil {
		return
	}
	traversal(node.Left)        //左
	res = append(res, node.Val) //中
	traversal(node.Right)       //右
}

func traversal(node *TreeNode) { //后序
	if node == nil {
		return
	}
	res = append(res, node.Val) //左
	traversal(node.Right)       //右
	traversal(node.Left)        //中
}


func invertTree(root *TreeNode) *TreeNode {  
    // var tran func(node *TreeNode)
    tran := func(node *TreeNode){
        if node==nil{
            return
        }
        node.Left,node.Right=node.Right,node.Left
        tran(node.Left)
		tran(node.Right)
    }
    tran(root)
    return root
}