package lessons

func relativeSortArray(arr1 []int, arr2 []int) []int {
	ans := make([]int, 0, len(arr1))
	count := make([]int, 1001)

	for i := 0; i < len(arr1); i++ {
		count[arr1[i]]++
	}

	for i := 0; i < len(arr2); i++ {
		for count[arr2[i]] > 0 {
			count[arr2[i]]--
			ans = append(ans, arr2[i])
		}
	}

	for val := 0; val <= 1000; val++ {
		for count[val] > 0 {
			count[val]--
			ans = append(ans, val)
		}
	}

	return ans
}
