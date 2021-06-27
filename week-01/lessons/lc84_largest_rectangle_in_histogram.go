package lessons

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	s := make([]Rect, 0)
	ans := 0

	for _, h := range heights {
		accumulatedWidth := 0
		for len(s) != 0 && s[len(s)-1].height >= h {
			accumulatedWidth += s[len(s)-1].width
			ans = max(ans, accumulatedWidth*s[len(s)-1].height)
			s = s[:len(s)-1]
		}
		s = append(s, Rect{h, accumulatedWidth + 1})
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

type Rect struct {
	height int
	width  int
}
