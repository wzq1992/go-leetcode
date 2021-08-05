package lessons

import "math"

func findMaxValueOfEquation(points [][]int, k int) int {
	ans := math.MinInt32
	var q []int
	for i := 0; i < len(points); i++ {
		// 求上界：j <= i - 1, 下界：x[j] >= x[i] - k
		// 这个范围中 y[j] - x[j] 的最大值
		// 考虑两个选项 j1 < j2
		// 写出 j1 比 j2 优的条件
		// y[j1] - x[j1] > y[j2] - x[j2]

		// 1. 队头合法性
		for len(q) > 0 && points[q[0]][0] < points[i][0]-k {
			q = q[1:]
		}
		// 2. 取队头为最优解
		if len(q) > 0 {
			ans = max(ans, points[i][1]+points[i][0]+points[q[0]][1]-points[q[0]][0])
		}
		// 3. 维护队列单调性，队尾插入新选项
		for len(q) > 0 && points[q[len(q)-1]][1]-points[q[len(q)-1]][0] <= points[i][1]-points[i][0] {
			q = q[:len(q)-1]
		}

		q = append(q, i)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
