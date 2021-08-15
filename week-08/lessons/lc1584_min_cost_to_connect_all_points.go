package lessons

import "sort"

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	// 构造出边
	var edges [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			edges = append(edges, []int{i, j, absInt(points[i][0]-points[j][0]) + absInt(points[i][1]-points[j][1])})
		}
	}

	// 按照边权排序
	sort.Sort((IntEdges)(edges))

	fa := make([]int, n)
	// Kurskal 算法
	for i := 0; i < n; i++ {
		fa[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if x == fa[x] {
			return x
		}

		fa[x] = find(fa[x])
		return fa[x]
	}

	ans := 0
	for _, edge := range edges {
		x, y, z := edge[0], edge[1], edge[2]

		if find(x) != find(y) {
			ans += z
			fa[find(x)] = find(y)
		}
	}

	return ans
}

func absInt(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

type IntEdges [][]int

func (p IntEdges) Len() int {
	return len(p)
}

func (p IntEdges) Less(i, j int) bool {
	return p[i][2] < p[j][2]
}

func (p IntEdges) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
