package lessons

// 方法二
func moveZeroes(nums []int) {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k {
				nums[k], nums[i] = nums[i], nums[k]
			}
			k++
		}
	}
}

// 方法一
func moveZeroes1(nums []int) {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}

	for k < len(nums) {
		nums[k] = 0
		k++
	}

}
