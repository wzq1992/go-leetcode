package lessons

func levelOrder(root *Node) (res [][]int) {
	if root == nil {
		return
	}

	q := []*Node{root}
	for len(q) > 0 {
		qSize := len(q)
		var node *Node
		var depthVals []int
		for i := 0; i < qSize; i++ {
			node = q[0]
			q = q[1:]
			depthVals = append(depthVals, node.Val)
			for _, child := range node.Children {
				q = append(q, child)
			}
		}

		res = append(res, depthVals)
	}

	return
}
