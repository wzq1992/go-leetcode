package lessons

func findCircleNum(isConnected [][]int) int {
	var fa []int
	n := len(isConnected)
	for i := 0; i < n; i++ {
		fa = append(fa, i)
	}

	var find func(x int) int
	find = func(x int) int {
		if x == fa[x] {
			return x
		}

		fa[x] = find(fa[x])
		return fa[x]
	}

	var unionSet func(i, j int)
	unionSet = func(i, j int) {
		x, y := find(i), find(j)
		if x != y {
			fa[x] = y
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				unionSet(i, j)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if find(i) == i {
			ans++
		}
	}

	return ans
}
