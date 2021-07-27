package homework

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, n := range piles {
		r = max(r, n)
	}

	for l < r {
		m := (l + r) / 2

		if needHours(piles, m) <= h {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func needHours(piles []int, k int) int {
	h := 0
	for _, n := range piles {
		h += n / k
		if n%k > 0 {
			h++
		}
	}

	return h
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
