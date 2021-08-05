package lessons

func findWords(board [][]byte, words []string) []string {
	var ans []string

	dx, dy := []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)

	var dfs func(x, y int, curr *Node)
	dfs = func(x, y int, curr *Node) {
		ch := board[x][y]
		if _, ok := curr.child[ch]; !ok {
			return
		}
		fa := curr
		curr = fa.child[ch]
		if curr.word != "" {
			ans = append(ans, curr.word)
			curr.word = ""
		}
		if len(curr.child) == 0 {
			delete(curr.child, ch)
		}

		for k := 0; k < 4; k++ {
			nx, ny := x + dx[k], y + dy[k]
			if nx < 0 || ny < 0 || nx >= m || ny >= n || visited[nx][ny] {
				continue
			}
			visited[nx][ny] = true
			dfs(nx, ny, curr)
			visited[nx][ny] = false
		}
	}

	root := NewNode()
	for _, word := range words {
		func(word string) {
			curr := root
			for i := 0; i < len(word); i++ {
				ch := word[i]
				if _, ok := curr.child[ch]; !ok {
					curr.child[ch] = NewNode()
				}

				curr = curr.child[ch]
			}

			curr.word = word
		}(word)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := range visited {
				visited[k] = make([]bool, n)
			}

			visited[i][j] = true
			dfs(i, j, root)
		}
	}

	return ans
}

type Node struct {
	word string
	child map[byte]*Node
}

func NewNode() *Node {
	return &Node{
		child: make(map[byte]*Node, 0),
	}
}
