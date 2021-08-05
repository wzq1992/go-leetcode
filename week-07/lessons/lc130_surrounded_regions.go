package lessons

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	fa := make([]int, m*n+1)

	var num func(i, j int) int
	num = func(i, j int) int {
		return i*n + j
	}

	var find func(x int) int
	find = func(x int) int {
		if x == fa[x] {
			return x
		}

		fa[x] = find(fa[x])
		return fa[x]
	}

	dx, dy := []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fa[num(i, j)] = num(i, j)
		}
	}

	fa[m*n] = m * n;
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if ni < 0 || nj < 0 || ni >= m || nj >= n {
					fa[find(num(i, j))] = find(m * n)
				} else if board[ni][nj] == 'O' {
					fa[find(num(ni, nj))] = find(num(i, j))
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && find(num(i, j)) != find(m*n) {
				board[i][j] = 'X'
			}
		}
	}
}
