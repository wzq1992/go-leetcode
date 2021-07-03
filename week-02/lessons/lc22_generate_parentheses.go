package lessons

var mCache = map[int][]string{}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	if cache, ok := mCache[n]; ok {
		return cache
	}

	var ans []string
	for k := 1; k <= n; k++ {
		resA := generateParenthesis(k - 1)
		resB := generateParenthesis(n - k)

		for _, a := range resA {
			for _, b := range resB {
				ans = append(ans, "("+a+")"+b)
			}
		}
	}

	mCache[n] = ans
	return ans
}
