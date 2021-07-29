package lessons

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	text1, text2 = " "+text1, " "+text2
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i] == text2[j] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}

	return f[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
