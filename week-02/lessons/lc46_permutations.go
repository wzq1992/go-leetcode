package lessons

func permute(nums []int) [][]int {
	var ans [][]int
	var per []int
	n := len(nums)
	used := make([]bool, n)

	var find func(int)

	find = func(index int) {
		if index == n {
			// make a copy
			s := make([]int, len(per))
			copy(s, per)
			ans = append(ans, s)
			return
		}

		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				per = append(per, nums[i])
				find(index + 1)
				per = per[:len(per)-1]
				used[i] = false
			}
		}
	}

	find(0)
	return ans
}
