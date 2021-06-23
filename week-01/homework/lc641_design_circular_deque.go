package homework

type MyCircularDeque struct {
	data     []int
	size     int
	capacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (c *MyCircularDeque) InsertFront(value int) bool {
	if c.IsFull() {
		return false
	}

	var newData []int
	newData = append(newData, value)
	newData = append(newData, c.data...)
	c.data = newData
	c.size++

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (c *MyCircularDeque) InsertLast(value int) bool {
	if c.IsFull() {
		return false
	}

	c.data = append(c.data, value)
	c.size++

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (c *MyCircularDeque) DeleteFront() bool {
	if c.IsEmpty() {
		return false
	}

	c.data = c.data[1:]
	c.size--

	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (c *MyCircularDeque) DeleteLast() bool {
	if c.IsEmpty() {
		return false
	}

	c.data = c.data[:c.size-1]
	c.size--

	return true
}

/** Get the front item from the deque. */
func (c *MyCircularDeque) GetFront() int {
	if c.IsEmpty() {
		return -1
	}

	return c.data[0]
}

/** Get the last item from the deque. */
func (c *MyCircularDeque) GetRear() int {
	if c.IsEmpty() {
		return -1
	}

	return c.data[c.size-1]
}

/** Checks whether the circular deque is empty or not. */
func (c *MyCircularDeque) IsEmpty() bool {
	return c.size == 0
}

/** Checks whether the circular deque is full or not. */
func (c *MyCircularDeque) IsFull() bool {
	return c.size == c.capacity
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
