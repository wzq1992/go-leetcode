package lessons

func lengthOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, f[i])
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
