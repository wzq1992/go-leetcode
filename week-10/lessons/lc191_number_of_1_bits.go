package lessons

func hammingWeight(num uint32) int {
	count := 0

	for num > 0 {
		count++
		num -= num & -num
	}

	return count
}
