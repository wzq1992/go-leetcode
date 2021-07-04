package lessons

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	dx := [4]int{-1, 0, 0, 1}
	dy := [4]int{0, -1, 1, 0}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int)

	dfs = func(x, y int) {
		visited[x][y] = true

		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}

			if grid[nx][ny] == '1' && !visited[nx][ny] {
				dfs(nx, ny)
			}
		}
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(i, j)
				ans++
			}
		}
	}

	return ans
}
