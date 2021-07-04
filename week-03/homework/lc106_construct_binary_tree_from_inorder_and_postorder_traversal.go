package homework

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	var build func(inorder []int, postorder []int, l1, r1, l2, r2 int) *TreeNode

	// [l1, r1) [l2, r2)
	build = func(inorder []int, postorder []int, l1, r1, l2, r2 int) *TreeNode {
		if l1 >= r1 {
			return nil
		}

		root := &TreeNode{
			Val: postorder[r2-1],
		}

		mid := l1
		for inorder[mid] != root.Val {
			mid++
		}

		leftSize := mid - l1

		root.Left = build(inorder, postorder, l1, mid, l2, l2+leftSize)
		root.Right = build(inorder, postorder, mid+1, r1, l2+leftSize, r2-1)

		return root
	}

	return build(inorder, postorder, 0, len(inorder), 0, len(postorder))
}
