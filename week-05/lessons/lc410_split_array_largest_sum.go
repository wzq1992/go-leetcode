package lessons

func splitArray(nums []int, m int) int {
	// 上下界
	left, right := 0, 0
	for _, num := range nums {
		left = max(left, num)
		right += num
	}

	for left < right {
		mid := (left + right) / 2
		// 第一个是的判定问题 IsValid 得到 true 的位置
		if IsValid(nums, m, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

// 判定把 nums 分成 <= m 组，每组的和 <= T
func IsValid(nums []int, m, T int) bool {
	groupSum, groupCount := 0, 1

	for i := 0; i < len(nums); i++ {
		if groupSum+nums[i] <= T {
			groupSum += nums[i]
		} else {
			groupCount++
			groupSum = nums[i]
		}
	}

	return groupCount <= m
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
