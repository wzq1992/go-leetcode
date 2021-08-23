package lessons

import (
	"container/heap"
	"strings"
)

type aStar struct {
	g, h   int
	status string
}
type hp []aStar

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].g+h[i].h < h[j].g+h[j].h }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(aStar)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 曼哈顿距离
var dist = [6][6]int{
	{0, 1, 2, 1, 2, 3},
	{1, 0, 1, 2, 1, 2},
	{2, 1, 0, 3, 2, 1},
	{1, 2, 3, 0, 1, 2},
	{2, 1, 2, 1, 0, 1},
	{3, 2, 1, 2, 1, 0},
}

// 计算启发函数
func getH(status string) (ret int) {
	for i, ch := range status {
		if ch != '0' {
			ret += dist[i][ch-'1']
		}
	}
	return
}

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func slidingPuzzle(board [][]int) int {
	const target = "123450"

	s := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}
	start := string(s)
	if start == target {
		return 0
	}

	// 枚举 status 通过一次交换操作得到的状态
	get := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	h := hp{{0, getH(start), start}}
	seen := map[string]bool{start: true}
	for len(h) > 0 {
		node := heap.Pop(&h).(aStar)
		for _, nxt := range get(node.status) {
			if !seen[nxt] {
				if nxt == target {
					return node.g + 1
				}
				seen[nxt] = true
				heap.Push(&h, aStar{node.g + 1, getH(nxt), nxt})
			}
		}
	}
	return -1
}
