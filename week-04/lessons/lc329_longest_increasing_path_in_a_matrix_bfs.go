package lessons

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	edges := make([][]int, m*n)
	deg := make([]int, m*n)
	dist := make([]int, m*n)
	dx := [...]int{-1, 0, 0, 1}
	dy := [...]int{0, -1, 1, 0}

	var topsort func()
	topsort = func() {
		var q []int

		for i := 0; i < m*n; i++ {
			if deg[i] == 0 {
				q = append(q, i)
				dist[i] = 1
			}
		}

		for len(q) > 0 {
			x := q[0]
			q = q[1:]

			for _, y := range edges[x] {
				deg[y]--
				dist[y] = max(dist[y], dist[x]+1)
				if deg[y] == 0 {
					q = append(q, y)
				}
			}
		}
	}

	var valid func(i, j int) bool
	valid = func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	var addEdge func(u, v int)
	addEdge = func(u, v int) {
		edges[u] = append(edges[u], v)
		deg[v]++
	}

	var num func(i, j int) int
	num = func(i, j int) int {
		return i*n + j
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// flag := true
			for k := 0; k < 4; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if valid(ni, nj) && matrix[i][j] < matrix[ni][nj] {
					addEdge(num(i, j), num(ni, nj))
				}
			}
		}
	}

	topsort()
	ans := 0
	for i := 0; i < m*n; i++ {
		ans = max(ans, dist[i])
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
