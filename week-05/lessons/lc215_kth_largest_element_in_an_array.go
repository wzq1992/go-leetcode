package lessons

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().Unix()))

func findKthLargest(nums []int, k int) int {
	// 第 K 大，等于（排序后下标 0 开始）第 N - K 小
	return quickSort(nums, len(nums)-k, 0, len(nums)-1)
}

func quickSort(nums []int, k, l, r int) int {
	if l == r {
		return nums[l]
	}

	p := partition(nums, l, r)

	if p >= k {
		return quickSort(nums, k, l, p)
	} else {
		return quickSort(nums, k, p+1, r)
	}
}

func partition(nums []int, l, r int) int {
	pivot := l + rnd.Int()%(r-l+1)
	pivotVal := nums[pivot]

	for l <= r {
		for nums[l] < pivotVal {
			l++
		}

		for nums[r] > pivotVal {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	return r
}
