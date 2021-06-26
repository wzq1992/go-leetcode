package lessons

func numberOfSubarrays(nums []int, k int) int {
	ans := 0

	nums = append([]int{0}, nums...)
	s := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		s[i] = s[i-1] + nums[i]%2
	}

	count := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		count[s[i]] += 1
	}

	for i := 1; i < len(nums); i++ {
		if s[i]-k >= 0 {
			ans += count[s[i]-k]
		}
	}

	return ans
}
