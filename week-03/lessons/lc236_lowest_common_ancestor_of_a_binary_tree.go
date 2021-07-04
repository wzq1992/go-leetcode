package lessons

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	fa := make(map[int]*TreeNode, 0)

	var calcFather func(root *TreeNode)
	calcFather = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			fa[root.Left.Val] = root
			calcFather(root.Left)
		}

		if root.Right != nil {
			fa[root.Right.Val] = root
			calcFather(root.Right)
		}
	}

	calcFather(root)

	redNodes := map[int]bool{root.Val: true}
	// p 往上走，标红
	for p.Val != root.Val {
		redNodes[p.Val] = true
		p = fa[p.Val]
	}

	// q 往上走，遇到第一个红色时结束
	for {
		if _, ok := redNodes[q.Val]; ok {
			break
		}

		q = fa[q.Val]
	}

	return q
}
