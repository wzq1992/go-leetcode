package lessons

func combine(n int, k int) [][]int {
	var set []int
	var ans [][]int

	var findSubsets func(int)
	findSubsets = func(index int) {
		// 剪枝
		if len(set)+n-index+1 < k {
			return
		}

		if len(set) == k {
			// make a copy
			s := make([]int, len(set))
			copy(s, set)
			ans = append(ans, s)
			return
		}

		// 不选
		findSubsets(index + 1)
		// 选
		set = append(set, index)
		findSubsets(index + 1)
		// 还原现场
		set = set[:len(set)-1]
	}

	findSubsets(1)
	return ans
}
