package homework

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < n; j++ {
			if len(strs[j]) == i || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
