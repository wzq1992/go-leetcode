package lessons

// 单调栈
func trap(height []int) int {
	s := []Rectangle{{0, 0}}
	ans := 0

	for _, h := range height {
		w := 0

		for len(s) > 1 && s[len(s)-1].height <= h {
			w += s[len(s)-1].width
			bottom := s[len(s)-1].height
			s = s[:len(s)-1]
			ans += w * max(0, min(s[len(s)-1].height, h)-bottom)
		}

		s = append(s, Rectangle{h, w + 1})
	}

	return ans
}

type Rectangle struct {
	height int
	width  int
}

// 方法一：前缀/后缀最大值
/*
func trap(height []int) int {
	n := len(height)
	pre, suf := make([]int, n + 1), make([]int, n + 2)

	for i := 1; i <= n; i++ {
		pre[i] = max(pre[i - 1], height[i - 1])
	}

	for i := n; i > 0; i-- {
		suf[i] = max(suf[i + 1], height[i - 1])
	}

	ans := 0
	for i := 1; i <= n; i ++ {
		ans += max(0, min(pre[i - 1], suf[i + 1]) - height[i - 1])
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}*/
