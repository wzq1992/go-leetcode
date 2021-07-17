package lessons

func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] <= nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}
