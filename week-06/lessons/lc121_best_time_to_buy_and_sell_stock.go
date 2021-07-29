package lessons

import "math"

func maxProfit(prices []int) int {
	priceMin := math.MaxInt32
	profitMax := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < priceMin {
			priceMin = prices[i]
		} else if prices[i]-priceMin > profitMax {
			profitMax = prices[i] - priceMin
		}
	}

	return profitMax
}
