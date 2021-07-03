package lessons

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	md := math.MaxInt64
	if root.Left != nil {
		md = min(md, minDepth(root.Left))
	}

	if root.Right != nil {
		md = min(md, minDepth(root.Right))
	}

	return md + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
