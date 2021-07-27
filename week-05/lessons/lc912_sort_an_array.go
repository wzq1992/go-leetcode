package lessons

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().Unix()))

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)

	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot)
	quickSort(nums, pivot+1, r)
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
