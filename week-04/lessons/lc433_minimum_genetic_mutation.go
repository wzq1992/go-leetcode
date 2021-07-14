package lessons

func minMutation(start string, end string, bank []string) int {
	q := []string{start}
	depth := make(map[string]int, 0)
	depth[start] = 0
	gen := [...]byte{'A', 'C', 'G', 'T'}

	bankSet := make(map[string]bool, 0)
	for _, b := range bank {
		bankSet[b] = true
	}

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		// 24 条出边
		for i := 0; i < 8; i++ {
			for j := 0; j < 4; j++ {
				if s[i] == gen[j] {
					continue
				}

				nsb := []byte(s)
				nsb[i] = gen[j]
				ns := string(nsb)
				if _, ok := bankSet[ns]; !ok {
					continue
				}

				// s -> ns
				// depth 里面没有 contains 字符串 ns
				if _, ok := depth[ns]; !ok {
					depth[ns] = depth[s] + 1
					q = append(q, ns)
					if ns == end {
						return depth[ns]
					}
				}
			}
		}
	}

	return -1
}
