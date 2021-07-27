package homework

var ans, lo, hi int

func countRangeSum(nums []int, lower int, upper int) int {
	ans, lo, hi = 0, lower, upper
	preSums := make([]int, len(nums)+1)
	for i, num := range nums {
		preSums[i+1] = preSums[i] + num
	}

	mergeSort(preSums, 0, len(preSums)-1)

	return ans
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	calculate(nums, l, mid, r)

	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, mid, r int) {
	newNums := make([]int, r-l+1)
	i, j := l, mid+1

	for k := range newNums {
		if j > r || (i <= mid && nums[i] <= nums[j]) {
			newNums[k] = nums[i]
			i++
		} else {
			newNums[k] = nums[j]
			j++
		}
	}

	for k, v := range newNums {
		nums[l+k] = v
	}
}

func calculate(nums []int, l, mid, r int) {
	i, j := mid+1, mid+1

	for k := l; k <= mid; k++ {
		for i <= r && nums[i]-nums[k] < lo {
			i++
		}

		for j <= r && nums[j]-nums[k] <= hi {
			j++
		}

		ans += j - i
	}
}
