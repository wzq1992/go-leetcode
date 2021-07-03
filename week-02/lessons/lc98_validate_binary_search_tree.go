package lessons

import "math"

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
}
