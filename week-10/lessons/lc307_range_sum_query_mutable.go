package lessons

import "math"

type NumArray struct {
	data   []int
	blocks []int
	n      int // 元素总数
	b      int // 每组个数
	bn     int // 组数
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	if n == 0 {
		return NumArray{}
	}

	b := int(math.Sqrt(float64(n)))
	bn := n / b
	if n%b > 0 {
		bn += 1
	}

	data := make([]int, n)
	copy(data, nums)

	blocks := make([]int, bn)
	for i := 0; i < n; i++ {
		blocks[i/b] += nums[i]
	}

	return NumArray{
		n:      n,
		b:      b,
		bn:     bn,
		data:   data,
		blocks: blocks,
	}
}

func (n *NumArray) Update(i int, val int) {
	if i < 0 || i > n.n {
		return
	}

	b := i / n.b
	n.blocks[b] -= n.data[i]
	n.blocks[b] += val

	n.data[i] = val
}

func (n *NumArray) SumRange(x int, y int) int {
	if x < 0 || x >= n.n || y < 0 || y >= n.n || x > y {
		return 0
	}

	bStart, bEnd := x/n.b, y/n.b

	res := 0

	if bStart == bEnd {
		for i := x; i <= y; i++ {
			res += n.data[i]
		}

		return res
	}

	for i := x; i < (bStart+1)*n.b; i++ {
		res += n.data[i]
	}

	for b := bStart + 1; b < bEnd; b++ {
		res += n.blocks[b]
	}

	for i := bEnd * n.b; i <= y; i++ {
		res += n.data[i]
	}

	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
