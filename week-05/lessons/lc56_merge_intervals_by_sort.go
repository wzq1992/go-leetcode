package lessons

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Sort((IntPairSlice)(intervals))

	var ans [][]int
	left, far := -1, -1
	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		if start <= far {
			far = max(far, end)
		} else {
			if far >= 0 {
				ans = append(ans, []int{left, far})
			}
			left, far = start, end
		}
	}

	if far >= 0 {
		ans = append(ans, []int{left, far})
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
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
