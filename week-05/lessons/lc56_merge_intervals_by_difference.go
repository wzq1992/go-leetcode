package lessons

import "sort"

func merge(intervals [][]int) [][]int {
	var events [][]int
	for _, interval := range intervals {
		events = append(events, []int{interval[0], 1})
		events = append(events, []int{interval[1] + 1, -1}) // 如果这里不 +1, 排序第二关键字保证 1 在 -1 之前
	}

	sort.Sort((IntPairSlice)(events))

	var ans [][]int
	left, count := 0, 0
	for _, event := range events {
		// 加之前是 0，加之后是非 0
		if count == 0 {
			left = event[0] // 一个段的产生
		}
		count += event[1]

		// 非 0 变 0，一个段的结束
		if count == 0 {
			ans = append(ans, []int{left, event[0] - 1})
		}
	}

	return ans
}

type IntPairSlice [][]int

func (p IntPairSlice) Len() int {
	return len(p)
}

func (p IntPairSlice) Less(i, j int) bool {
	return p[i][0] < p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1])
}

func (p IntPairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
