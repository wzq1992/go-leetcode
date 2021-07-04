package lessons

// 迭代法
func preorder(root *Node) (res []int) {
	if root == nil {
		return
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)

		for i := len(root.Children) - 1; i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
	}

	return
}

// 递归法
func preorderRecursion(root *Node) (res []int) {
	var find func(root *Node)

	find = func(root *Node) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		for _, child := range root.Children {
			find(child)
		}
	}

	find(root)

	return res
}
