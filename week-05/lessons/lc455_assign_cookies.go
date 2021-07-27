package lessons

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	j, ans := 0, 0
	for i := 0; i < len(g); i++ {
		for j < len(s) && g[i] > s[j] {
			j++
		}

		if j < len(s) {
			j++
			ans++
		}
	}

	return ans
}
