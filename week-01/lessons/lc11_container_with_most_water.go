package lessons

func maxArea(height []int) int {
	i, j := 0, len(height)-1

	ans := 0
	for i < j {
		ans = max1(min1(height[i], height[j])*(j-i), ans)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return ans
}

func max1(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min1(x, y int) int {
	if x < y {
		return x
	}

	return y
}
