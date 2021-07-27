package homework

func shipWithinDays(weights []int, days int) int {
	var l, r int
	for _, w := range weights {
		l = max(l, w)
		r += w
	}

	for l < r {
		mid := (l + r) / 2

		if needDays(weights, mid) <= days {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func needDays(weights []int, k int) int {
	res, sum := 1, 0
	for _, w := range weights {
		if sum+w > k {
			res++
			sum = 0
		}
		sum += w
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
