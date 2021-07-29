package lessons

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+1)
	word1 = " " + word1
	word2 = " " + word2

	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
		f[i][0] = i
	}
	for i := 0; i <= n; i++ {
		f[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			eq := 0
			if word1[i] != word2[j] {
				eq = 1
			}
			f[i][j] = min(min(f[i-1][j], f[i][j-1])+1, f[i-1][j-1]+eq)
		}
	}

	return f[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
