package lessons

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	edges := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var ans []string
	var s []byte
	var dfs func(string, int)

	dfs = func(digits string, index int) {
		// 终止条件
		if len(digits) == index {
			ans = append(ans, string(s))
			return
		}

		// 考虑所有出边
		for _, ch := range edges[digits[index]] {
			s = append(s, ch)
			dfs(digits, index+1)
			s = s[:len(s)-1]
		}
	}

	dfs(digits, 0)

	return ans
}
