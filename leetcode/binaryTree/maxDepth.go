//二叉树最大深度

func maxDepth(root *TreeNode) int { //迭代法
	if root == nil {
		return 0
	}
	max := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		len := len(queue)
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

func maxDepth(root *TreeNode) int {
	var getd func(node *TreeNode) int
	getd = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ld := getd(node.Left)
		rd := getd(node.Right)
		dep := 1 + max(ld, rd)
		return dep
	}
	return getd(root)
}

func maxDepth(root *TreeNode) int {
	var getd func(node *TreeNode) int
	getd = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ld := getd(node.Left)
		rd := getd(node.Right)
		max := func(ld, rd int) int {
			if ld > rd {
				return ld
			} else {
				return rd
			}
		}(ld, rd)
		return 1 + max
	}
	return getd(root)
}
