package lessons

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return nil
}
