package lessons

import "math"

func coinChange(coins []int, amount int) int {
	opt := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		opt[i] = math.MaxInt32
		for _, c := range coins {
			if i-c >= 0 {
				opt[i] = min(opt[i], opt[i-c]+1)
			}
		}
	}

	if opt[amount] >= math.MaxInt32 {
		opt[amount] = -1
	}

	return opt[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
