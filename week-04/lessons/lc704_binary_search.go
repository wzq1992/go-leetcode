package lessons

func search(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2

		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			l = mid + 1
			continue
		}

		r = mid
	}

	return -1
}
