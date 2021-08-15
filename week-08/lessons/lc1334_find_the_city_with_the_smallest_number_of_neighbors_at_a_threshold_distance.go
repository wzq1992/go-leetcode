package lessons

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = math.MaxInt32
		}
		d[i][i] = 0
	}

	for _, edge := range edges {
		x, y, z := edge[0], edge[1], edge[2]
		d[x][y] = z
		d[y][x] = z
	}

	// Floyd 算法
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	ansCount, ans := n, 0
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j ++ {
			if i != j && d[i][j] <= distanceThreshold {
				count++
			}
		}

		if count <= ansCount {
			ansCount = count
			ans = i
		}
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
