package lessons

func treeDiameter(edges [][]int) int {
	if edges == nil || len(edges) == 0 {
		return 0
	}
	tree := make(map[int][]int, len(edges))
	for _, v := range edges {
		v0, v1 := v[0], v[1]
		if v1 < v0 {
			v0, v1 = v[1], v[0]
		}
		nodes, ok := tree[v0]
		if !ok {
			nodes = []int{}
		}
		tree[v0] = append(nodes, v1)
	}
	_, path := dfs(tree, 0)
	return path
}

func dfs(tree map[int][]int, k int) (int, int) {
	childs, ok := tree[k]
	if !ok || len(childs) == 0 {
		return 0, 0
	}
	diameter, h1, h2 := 0, -1, -1
	for _, c := range childs {
		h, d := dfs(tree, c)
		if diameter < d {
			diameter = d
		}
		if h1 < h {
			h1, h2 = h, h1
		} else if h2 < h {
			h2 = h
		}
	}
	if diameter < h1+h2+2 {
		diameter = h1 + h2 + 2
	}
	return h1 + 1, diameter
}
