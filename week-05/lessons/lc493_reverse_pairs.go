package lessons

var ans int

func reversePairs(nums []int) int {
	ans = 0
	mergeSort(nums, 0, len(nums)-1)
	return ans
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	calculate(nums, l, mid, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, left, mid, right int) {
	newNums := make([]int, right-left+1)
	i, j := left, mid+1

	for k := 0; k < len(newNums); k++ {
		if j > right || (i <= mid && nums[i] <= nums[j]) {
			newNums[k] = nums[i]
			i++
		} else {
			newNums[k] = nums[j]
			j++
		}
	}

	for k := 0; k < len(newNums); k++ {
		nums[left+k] = newNums[k]
	}
}

func calculate(nums []int, left, mid, right int) {
	for i, j := left, mid; i <= mid; i++ {
		for j < right && nums[i] > 2*nums[j+1] {
			j++
		}
		ans += j - mid
	}
}
