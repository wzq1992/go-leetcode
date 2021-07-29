package lessons

func canPartition(nums []int) bool {
	n := len(nums)
	nums = append([]int{0}, nums...)

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2
	f := make([]bool, sum+1)
	f[0] = true

	for i := 1; i <= n; i++ {
		for j := sum; j >= nums[i]; j-- {
			f[j] = f[j] || f[j-nums[i]]
		}
	}

	return f[sum]
}
