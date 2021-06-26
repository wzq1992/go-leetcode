package lessons

import "math"

type MinStack struct {
	preMin []int
	stack  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		preMin: []int{math.MaxInt64},
		stack:  []int{},
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	m.preMin = append(m.preMin, min(m.preMin[len(m.preMin)-1], val))
}

func (m *MinStack) Pop() {
	m.preMin = m.preMin[:len(m.preMin)-1]
	m.stack = m.stack[:len(m.stack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.preMin[len(m.preMin)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
