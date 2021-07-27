package lessons

func jump(nums []int) int {
	now, ans := 0, 0

	for now < len(nums)-1 {
		if nums[now] == 0 {
			return -1
		}

		right := now + nums[now]
		if right >= len(nums)-1 {
			return ans + 1
		}

		next := now + 1
		for i := now + 2; i <= right; i++ {
			nextRight := i + nums[i]
			if nextRight > next+nums[next] {
				next = i
			}
		}

		now = next
		ans++
	}

	return ans
}
