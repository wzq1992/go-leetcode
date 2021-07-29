package lessons

func maxProfit(prices []int) int {
	n := len(prices)

	// 只保留上一次状态
	f := []int{0, -prices[0]}
	for i := 1; i < n; i++ {
		f[0] = max(f[0], f[1]+prices[i])
		f[1] = max(f[1], f[0]-prices[i])
	}

	return f[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
