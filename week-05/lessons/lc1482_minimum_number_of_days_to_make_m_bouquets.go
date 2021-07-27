package lessons

func minDays(bloomDay []int, m int, k int) int {
	maxValue := 1000000001
	left, right := 0, maxValue

	for left < right {
		mid := (left + right) / 2
		if isValidOnDay(bloomDay, m, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left == maxValue {
		return -1
	}

	return left
}

func isValidOnDay(bloomDay []int, m, k, T int) bool {
	consecutive := 0 // 已经连续开了几朵
	bouquet := 0     // 花束数量

	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] <= T { // 花开了
			consecutive++
			if consecutive == k {
				bouquet++
				consecutive = 0
			}
		} else {
			// 花没开
			consecutive = 0
		}
	}

	return bouquet >= m
}
