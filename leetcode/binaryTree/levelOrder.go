//102二叉树层序遍历
package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *TreeNode) [][]int { //二叉树层序遍历
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmp []int
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层长度
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val) //将值加入本层切片
		}
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}

func levelOrderN(root *Node) [][]int { //n叉数层序遍历
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmp []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Children != nil {
				for _, i := range node.Children {
					queue.PushBack(i)
				}
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}

func levelOrderBottom(root *TreeNode) [][]int { //层序遍历加一个反转
	res := [][]int{}
	var tmp []int
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		if tmp[0] != 0 {
			res = append(res, tmp)
		}
		tmp = []int{}
	}
	for i, j := 0, len(res)-1; i < j; i++ {
		res[i], res[j] = res[j], res[i]
		j--
	}
	return res
}
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmp []int
	height := 0
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层长度
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			// if node.Left != nil {
			// 	queue.PushBack(node.Left)
			// }
			if node.Right != nil {
				queue.PushBack(node.Right)
				height += 1
			}
			tmp = append(tmp, node.Val) //将值加入本层切片
		}
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	var traversal func(node *TreeNode, h int)
	t := [105]int{}
	traversal = func(node *TreeNode, h int) {
		if node == nil {
			return
		}
		if t[h] == 0 {
			res = append(res, node.Val)
			t[h] = 1
		}
		traversal(node.Right, h+1)
		traversal(node.Left, h+1)
	}
	traversal(root, 1)
	return res
}

func rightSideView(root *TreeNode) [][]int { //二叉树层序遍历
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmp []int
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层长度
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Left == nil {
				tmp = append(tmp, node.Val) //将值加入本层切片
			}

		}
		res = append(res, tmp)
		tmp = []int{}
	}
	return res
}

func maxDepth(root *TreeNode) int { //二叉树最大深度
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	max := 0
	for len(queue) > 0 {
		len := len(queue) // 这里一定要使用固定大小size，队列大小一直变化
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		max += 1
	}
	return max
}
