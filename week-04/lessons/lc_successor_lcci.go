package lessons

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return findSucc(root, p.Val)
}

func findSucc(root *TreeNode, val int) *TreeNode {
	curr := root
	var ans *TreeNode

	for curr != nil {
		if curr.Val > val && (ans == nil || ans.Val > curr.Val) {
			ans = curr
		}

		if curr.Val == val && curr.Right != nil {
			curr = curr.Right

			for curr.Left != nil {
				curr = curr.Left
			}

			return curr
		}

		if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return ans
}
