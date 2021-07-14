package lessons

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	answer := make([][]int, m)
	for i := 0; i < m; i++ {
		answer[i] = make([]int, n)
		for j := 0; j < n; j++ {
			answer[i][j] = -1
		}
	}

	dx := [...]int{-1, 0, 0, 1}
	dy := [...]int{0, -1, 1, 0}

	var howFar func(matrix [][]int, x, y int) int

	howFar = func(matrix [][]int, x, y int) int {
		if answer[x][y] != -1 {
			return answer[x][y]
		}

		answer[x][y] = 1

		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}
			if matrix[nx][ny] > matrix[x][y] {
				answer[x][y] = max(answer[x][y], howFar(matrix, nx, ny)+1)
			}
		}

		return answer[x][y]
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, howFar(matrix, i, j))
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
