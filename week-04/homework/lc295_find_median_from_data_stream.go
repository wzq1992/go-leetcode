package homework

import "container/heap"

type MedianFinder struct {
	mins *MaxIntHeap
	maxs *MinIntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {

	medianFinder := MedianFinder{
		mins: &MaxIntHeap{},
		maxs: &MinIntHeap{},
	}

	heap.Init(medianFinder.mins)
	heap.Init(medianFinder.maxs)

	return medianFinder
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxs, num)
	heap.Push(this.mins, heap.Pop(this.maxs))
	if this.mins.Len() > this.maxs.Len() {
		heap.Push(this.maxs, heap.Pop(this.mins))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	var maxsTop, minsTop int
	maxsTop = heap.Pop(this.maxs).(int)
	heap.Push(this.maxs, maxsTop)

	if (this.mins.Len()+this.maxs.Len())%2 == 1 {
		return float64(maxsTop)
	}

	minsTop = heap.Pop(this.mins).(int)
	heap.Push(this.mins, minsTop)

	return float64(minsTop+maxsTop) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 大堆
type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
