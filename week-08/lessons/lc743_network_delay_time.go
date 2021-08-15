package lessons

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32 / 2
	}

	dist[k] = 0
	// Bellman-ford
	for r := 1; r < n; r++ {
		flag := false
		for _, edge := range times {
			x, y, z := edge[0], edge[1], edge[2]
			if dist[x]+z < dist[y] {
				dist[y] = dist[x] + z
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	ans := 0
	for i := 0; i <= n; i++ {
		ans = max(ans, dist[i])
	}

	if ans == math.MaxInt32/2 {
		ans = -1
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
