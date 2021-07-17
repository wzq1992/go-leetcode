package homework

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right --
		}
	}

	return nums[left]
}
