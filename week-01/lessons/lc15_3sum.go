package lessons

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		allTwoSums := twoSum1(nums, i+1, -nums[i])
		for _, jk := range allTwoSums {
			ans = append(ans, []int{nums[i], jk[0], jk[1]})
		}
	}

	return ans
}

func twoSum1(nums []int, start int, target int) [][]int {
	ans := make([][]int, 0)
	i, j := start, len(nums)-1
	for i < j {
		if i > start && nums[i] == nums[i-1] {
			i++
			continue
		}

		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			ans = append(ans, []int{nums[i], nums[j]})
			i++
		}

	}
	return ans
}
