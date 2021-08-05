package homework

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	fa := make([]int, m*n)
	for i := range fa {
		fa[i] = i
	}

	var find func(x int) int
	var union func(i, j int) bool
	var getIndex func(i, j int) int

	find = func(x int) int {
		if x == fa[x] {
			return x
		}

		fa[x] = find(fa[x])
		return fa[x]
	}

	union = func(i, j int) bool {
		x, y := find(i), find(j)
		if x == y {
			return false
		}

		fa[x] = y
		return true
	}

	getIndex = func(i, j int) int {
		return i*n + j
	}

	ans := 0
	dx, dy := []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				for k := 0; k < 4; k++ {
					ni, nj := i+dx[k], j+dy[k]
					if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == '1' {
						if union(getIndex(i, j), getIndex(ni, nj)) {
							ans--
						}
					}
				}
				grid[i][j] = '0'
			}
		}
	}

	return ans
}
