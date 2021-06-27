package lessons

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	q := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(q) != 0 && q[0] <= i-k {
			q = q[1:]
		}

		for len(q) != 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}

		q = append(q, i)
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}
