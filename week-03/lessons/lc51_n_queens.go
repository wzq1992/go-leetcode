package lessons

func solveNQueens(n int) [][]string {
	var find func(int)
	var ans [][]int
	var s []int

	used := make([]bool, n)
	usedIplusJ := make(map[int]bool, n)
	usedIminusJ := make(map[int]bool, n)

	find = func(row int) {
		if row == n {
			// make a copy
			ss := make([]int, len(s))
			copy(ss, s)
			ans = append(ans, ss)

			return
		}

		for col := 0; col < n; col++ {
			if !used[col] && !usedIplusJ[row+col] && !usedIminusJ[row-col] {
				used[col] = true
				usedIplusJ[row+col] = true
				usedIminusJ[row-col] = true
				s = append(s, col)
				find(row + 1)
				s = s[:len(s)-1]
				usedIminusJ[row-col] = false
				usedIplusJ[row+col] = false
				used[col] = false
			}
		}
	}

	find(0)

	var result [][]string
	for _, per := range ans {
		var res []string
		for row := 0; row < n; row++ {
			s := make([]byte, n)
			for col := 0; col < n; col++ {
				if per[row] == col {
					s[col] = 'Q'
				} else {
					s[col] = '.'
				}
			}

			res = append(res, string(s))
		}

		result = append(result, res)
	}

	return result
}
