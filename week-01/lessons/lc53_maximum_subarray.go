package lessons

import "math"

func maxSubArray(nums []int) int {
	sum := make([]int, len(nums)+1)
	for i := 1; i < len(nums)+1; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	ans := math.MinInt64
	preMin := sum[0]
	for i := 1; i < len(nums)+1; i++ {
		ans = max1(ans, sum[i]-preMin)
		preMin = min1(preMin, sum[i])
	}

	return ans
}

func max1(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min1(x, y int) int {
	if x < y {
		return x
	}

	return y
}
