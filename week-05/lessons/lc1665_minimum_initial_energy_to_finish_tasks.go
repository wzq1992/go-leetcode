package lessons

import "sort"

func minimumEffort(tasks [][]int) int {
	// 消耗能量小，门槛高是现做的条件
	sort.Sort((IntPairSlice)(tasks))

	// 正序做任务，但计算要倒序
	energy := 0 // 任务全部做完的时候，还需要 0 的能量
	for i := len(tasks) - 1; i >= 0; i-- {
		energy = max(tasks[i][1], energy+tasks[i][0])
	}

	return energy
}

type IntPairSlice [][]int

func (p IntPairSlice) Len() int {
	return len(p)
}

func (p IntPairSlice) Less(i, j int) bool {
	return p[i][0]-p[i][1] < p[j][0]-p[j][1]
}

func (p IntPairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
