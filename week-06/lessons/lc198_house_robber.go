package lessons

import "math"

func rob(nums []int) int {
	n := len(nums)
	f := make([][2]int, n+1)
	nums = append([]int{0}, nums...)

	f[0][0], f[0][1] = 0, math.MinInt32

	for i := 1; i <= n; i++ {
		f[i][0] = max(f[i-1][1], f[i-1][0])
		f[i][1] = f[i-1][0] + nums[i]
	}

	return max(f[n][0], f[n][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
