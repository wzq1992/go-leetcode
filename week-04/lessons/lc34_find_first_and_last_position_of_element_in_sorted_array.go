package lessons

func searchRange(nums []int, target int) []int {
	left := getRight(nums, target-1)
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	return []int{left, getRight(nums, target) - 1}
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
