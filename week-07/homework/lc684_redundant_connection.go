package homework

func findRedundantConnection(input [][]int) []int {
	n := len(input)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}

	var find func(x int) int
	var union func(i, j int) bool

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

	for _, edge := range input {
		i, j := edge[0], edge[1]
		if !union(i, j) {
			return edge
		}
	}

	return nil
}
