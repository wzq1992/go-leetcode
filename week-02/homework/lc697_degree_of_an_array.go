package homework

func findShortestSubArray(nums []int) int {
	m := map[int]Item{}
	maxCount := 0
	for i, n := range nums {
		item, ok := m[n]
		if !ok {
			item = Item{1, i, i}
		} else {
			item.count++
			item.last = i
		}
		m[n] = item

		if item.count > maxCount {
			maxCount = item.count
		}
	}

	minLen := len(nums)
	for _, item := range m {
		length := item.last - item.first + 1
		if item.count == maxCount && length < minLen {
			minLen = length
		}
	}

	return minLen
}

type Item struct {
	count int
	first int
	last  int
}
