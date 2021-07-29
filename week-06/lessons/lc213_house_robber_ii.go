package lessons

import "math"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	f := make([][2]int, n+1)
	nums = append([]int{0}, nums...)

	// 不偷 1
	f[1][0], f[1][1] = 0, math.MinInt32
	for i := 2; i <= n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1])
		f[i][1] = f[i-1][0] + nums[i]
	}
	ans := max(f[n][0], f[n][1])

	// 不偷 n
	f[1][0], f[1][1] = 0, nums[1]
	for i := 2; i <= n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1])
		f[i][1] = f[i-1][0] + nums[i]
	}

	return max(ans, f[n][0])
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
