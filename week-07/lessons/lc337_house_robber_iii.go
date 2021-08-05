package lessons

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var dp = make(map[*TreeNode][]int, 0)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dp[root] = []int{0, root.Val}

	if root.Left != nil {
		dp[root][0] += rob(root.Left)
		dp[root][1] += dp[root.Left][0]
	}

	if root.Right != nil {
		dp[root][0] += rob(root.Right)
		dp[root][1] += dp[root.Right][0]
	}

	return max(dp[root][0], dp[root][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
