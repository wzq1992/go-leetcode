package lessons

import "math"

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	nums = append([]int{0}, nums...)

	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	// 最大子段和
	ans := math.MinInt32
	temp := preSum[0]
	for i := 1; i <= n; i++ {
		temp = min(temp, preSum[i-1])
		ans = max(ans, preSum[i]-temp)
	}

	// 最小子段和
	ansMin := math.MaxInt32
	temp = preSum[0]
	for i := 1; i < n; i++ {
		temp = max(temp, preSum[i-1])
		ansMin = min(ansMin, preSum[i]-temp)
	}

	temp = preSum[1]
	for i := 2; i < n; i++ {
		temp = max(temp, preSum[i])
	}

	if n > 1 {
		ansMin = min(ansMin, preSum[n]-temp)
	}

	return max(ans, preSum[n]-ansMin)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
