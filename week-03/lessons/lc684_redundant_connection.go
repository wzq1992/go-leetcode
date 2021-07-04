package lessons

func findRedundantConnection(input [][]int) []int {
	var addEdge func(x, y int)
	var dfs func(x, fa int)

	n := len(input)

	edges := make([][]int, n+1)
	visit := make([]bool, n+1)
	hasCycle := false

	addEdge = func(x, y int) {
		edges[x] = append(edges[x], y)
	}

	dfs = func(x, fa int) {
		// 第一步: 标记已访问
		visit[x] = true

		// 第二步: 遍历所有出边
		for _, y := range edges[x] {
			if y == fa {
				continue
			}

			if visit[y] {
				hasCycle = true
			} else {
				dfs(y, x)
			}
		}

		visit[x] = false
	}

	// 加边
	for _, edge := range input {
		u, v := edge[0], edge[1]

		//  无向图看做是双向边的有向图
		addEdge(u, v)
		addEdge(v, u)

		// 每加一条边就判断是否有环
		dfs(u, -1)
		if hasCycle {
			return edge
		}
	}

	return nil
}

