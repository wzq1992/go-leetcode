package homework

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	// 方向数组
	dx := [4]int{-1, 0, 0, 1}
	dy := [4]int{0, -1, 1, 0}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int)

	dfs = func(x, y int) {
		if board[x][y] != 'O' {
			return
		}
		visited[x][y] = true

		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}

			if !visited[nx][ny] {
				dfs(nx, ny)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
