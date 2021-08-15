package lessons

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	s, p = " "+s, " "+p
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true

	for i := 2; i <= n && p[i] == '*'; i += 2 {
		f[0][i] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j] == '.' {
				f[i][j] = f[i-1][j-1]
			} else if p[j] == '*' {
				f[i][j] = f[i][j-2]
				if p[j-1] == '.' || s[i] == p[j-1] {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else {
				f[i][j] = f[i-1][j-1] && s[i] == p[j]
			}
		}
	}

	return f[m][n]
}
