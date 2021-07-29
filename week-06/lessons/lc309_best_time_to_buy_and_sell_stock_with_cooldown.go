package lessons

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	f := make([][2][2]int, n+1)
	prices = append([]int{0}, prices...)
	for i := range f {
		f[i][0][0] = math.MinInt32
		f[i][0][1] = math.MinInt32
		f[i][1][0] = math.MinInt32
		f[i][1][1] = math.MinInt32
	}
	f[0][0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= 1; j++ {
			for k := 0; k <= 1; k++ {
				if j == 0 && k == 0 {
					f[i+1][1][0] = max(f[i+1][1][0], f[i][0][k]-prices[i+1])
				}

				if j == 1 && k == 0 {
					f[i+1][0][1] = max(f[i+1][0][1], f[i][1][k]+prices[i+1])
				}

				f[i+1][j][0] = max(f[i+1][j][0], f[i][j][k])
			}
		}
	}

	ans := 0
	for i := 0; i <= 1; i++ {
		ans = max(ans, f[n][0][i])
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
