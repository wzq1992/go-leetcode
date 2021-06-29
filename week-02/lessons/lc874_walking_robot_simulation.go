package lessons

func robotSim(commands []int, obstacles [][]int) int {
	blockers := map[int64]bool{}
	for _, obstacle := range obstacles {
		blockers[hash(obstacle[0], obstacle[1])] = true
	}

	// 方向数组	   N  E  S   W
	dx := [...]int{0, 1, 0, -1}
	dy := [...]int{1, 0, -1, 0}
	x, y := 0, 0
	dir := 0

	ans := 0
	for _, cmd := range commands {
		if cmd > 0 {
			for i := 0; i < cmd; i++ {
				nextX := x + dx[dir]
				nextY := y + dy[dir]

				if _, ok := blockers[hash(nextX, nextY)]; ok {
					break
				}

				x = nextX
				y = nextY
				ans = max(ans, x*x+y*y)
			}
		} else if cmd == -1 {
			// 右转
			dir = (dir + 1) % 4
		} else {
			// 左转
			dir = (dir - 1 + 4) % 4
		}
	}

	return ans
}

// func hash(x, y int) string {
// 	return strconv.Itoa(x) + "," + strconv.Itoa(y)
// }

func hash(x, y int) int64 {
	return (int64(x)+30000)*60000 + int64(y) + 30000
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
