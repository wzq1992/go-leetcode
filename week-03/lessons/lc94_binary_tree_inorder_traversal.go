package lessons

func inorderTraversal(root *TreeNode) (res []int) {
	var find func(root *TreeNode)

	find = func(root *TreeNode) {
		if root == nil {
			return
		}
		find(root.Left)
		res = append(res, root.Val)
		find(root.Right)
	}

	find(root)

	return res
}

// 迭代法
func inorderTraversalIteration(root *TreeNode) (res []int) {
	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}

	return
}
