package homework

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1

	for _, n := range nums {
		pre += n

		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}

		m[pre] += 1
	}

	return count
}
