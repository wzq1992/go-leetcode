package lessons

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	prices = append([]int{0}, prices...)

	f := make([][2][]int, n+1)
	for i := range f {
		for j := 0; j <= k; j++ {
			f[i][0] = append(f[i][0], math.MinInt32)
			f[i][1] = append(f[i][1], math.MinInt32)
		}
	}

	f[0][0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= 1; j++ {
			for c := 0; c <= k; c++ {
				f[i][j][c] = f[i-1][j][c]

				if j == 0 {
					f[i][0][c] = max(f[i][0][c], f[i-1][1][c]+prices[i])
				}
				if j == 1 && c > 0 {
					f[i][1][c] = max(f[i][1][c], f[i-1][0][c-1]-prices[i])
				}
			}
		}
	}

	ans := 0
	for i := 0; i <= k; i++ {
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
