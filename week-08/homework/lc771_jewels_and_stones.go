package homework

func numJewelsInStones(jewels string, stones string) int {
	m, n := len(jewels), len(stones)
	h := make(map[byte]bool, m)

	for i := 0; i < m; i++ {
		h[jewels[i]] = true
	}

	ans := 0
	for i := 0; i < n; i++ {
		if _, ok := h[stones[i]]; ok {
			ans++
		}
	}

	return ans
}
