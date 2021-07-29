package lessons

func maxProduct(nums []int) int {
	n := len(nums)
	fMin := make([]int, n)
	fMax := make([]int, n)

	fMin[0], fMax[0] = nums[0], nums[0]
	ans := fMax[0]
	for i := 1; i < n; i++ {
		fMax[i] = max(max(fMax[i-1]*nums[i], fMin[i-1]*nums[i]), nums[i])
		fMin[i] = min(min(fMax[i-1]*nums[i], fMin[i-1]*nums[i]), nums[i])

		ans = max(ans, fMax[i])
	}

	return ans
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
