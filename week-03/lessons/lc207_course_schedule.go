package lessons

func canFinish(numCourses int, prerequisites [][]int) bool {
	n := numCourses
	edges := make([][]int, n)
	inDeg := make([]int, n)
	var addEdge func(int, int)
	var topsort func() int
	var learned int

	addEdge = func(x, y int) {
		edges[x] = append(edges[x], y)
		inDeg[y]++
	}

	topsort = func() int {
		// 拓扑排序基于 BSF，需要队列
		var q []int
		// 从所有 0 入度点出发
		for i, deg := range inDeg {
			if deg == 0 {
				q = append(q, i)
			}
		}
		// BFS
		for len(q) != 0 {
			x := q[0] // 队头，这门课学了
			q = q[1:]
			learned++

			// 考虑 x 的所有边
			for _, y := range edges[x] {
				inDeg[y]-- // 去掉约束关系
				if inDeg[y] == 0 {
					q = append(q, y)
				}
			}
		}

		return learned
	}

	for _, pre := range prerequisites {
		addEdge(pre[0], pre[1])
	}

	return topsort() == n
}
