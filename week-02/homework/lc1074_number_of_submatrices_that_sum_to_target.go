package homework

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	ans := 0
	for i := range matrix {
		nums := make([]int, len(matrix[i]))

		// 依次计算 [i, i], [i, i+1], ... , [i, len(matrix)) 行每一列的和
		// 然后求和为 target 的子数组的个数
		for _, row := range matrix[i:] {
			for c, v := range row {
				nums[c] += v
			}

			ans += subArraySum(nums, target)
		}

	}

	return ans
}

// lc560
func subArraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{0: 1}

	for _, n := range nums {
		pre += n

		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}

		m[pre]++
	}

	return count
}
