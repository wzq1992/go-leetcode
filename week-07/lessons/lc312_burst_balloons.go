package lessons

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	f := make([][]int, n+2)
	for i := range f {
		f[i] = make([]int, n+2)
	}

	// 区间 DP，先枚举区间长度
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			r := i + l - 1
			for p := i; p <= r; p++ {
				f[i][r] = max(f[i][r], f[i][p-1]+f[p+1][r]+nums[i-1]*nums[p]*nums[r+1])
			}
		}
	}

	return f[1][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
