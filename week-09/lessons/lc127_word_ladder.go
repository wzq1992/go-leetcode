package lessons

func ladderLength(beginWord string, endWord string, wordList []string) int {
	distBegin, distEnd := make(map[string]int, 0), make(map[string]int, 0)

	for _, word := range wordList {
		distBegin[word], distEnd[word] = int(1e9), int(1e9)
	}

	// endWord 不在字典中
	if _, ok := distBegin[endWord]; !ok {
		return 0
	}

	distBegin[beginWord], distEnd[endWord] = 1, 1
	qBegin, qEnd := []string{beginWord}, []string{endWord}

	for len(qBegin) > 0 || len(qEnd) > 0 {
		var res int
		if qBegin, res = expand(qBegin, distBegin, distEnd); res != -1 {
			return res
		}

		if qEnd, res = expand(qEnd, distEnd, distBegin); res != -1 {
			return res
		}
	}

	return 0
}

func expand(q []string, dist map[string]int, distOther map[string]int) ([]string, int) {
	if len(q) > 0 {
		s := []rune(q[0])
		q = q[1:]
		depth := dist[string(s)]

		for i := 0; i < len(s); i++ {
			for ch := 'a'; ch <= 'z'; ch++ {
				backup := s[i]
				s[i] = ch

				if _, ok := dist[string(s)]; ok && dist[string(s)] > depth+1 {
					dist[string(s)] = depth + 1
					q = append(q, string(s))

					if distOther[string(s)] != int(1e9) {
						return q, depth + distOther[string(s)]
					}
				}

				s[i] = backup
			}
		}
	}
	return q, -1
}
