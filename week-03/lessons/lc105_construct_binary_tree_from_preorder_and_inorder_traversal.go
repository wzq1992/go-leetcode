package lessons

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var build func(preorder []int, inorder []int, l1, r1, l2, r2 int) *TreeNode

	// [l1, r1) [l2, r2)
	build = func(preorder []int, inorder []int, l1, r1, l2, r2 int) *TreeNode {
		if l1 >= r1 {
			return nil
		}

		root := &TreeNode{
			Val: preorder[l1],
		}

		mid := l2
		for inorder[mid] != root.Val {
			mid++
		}

		leftSize := mid - l2

		root.Left = build(preorder, inorder, l1+1, l1+1+leftSize, l2, mid)
		root.Right = build(preorder, inorder, l1+1+leftSize, r1, mid+1, r2)

		return root
	}

	return build(preorder, inorder, 0, len(preorder), 0, len(inorder))
}
