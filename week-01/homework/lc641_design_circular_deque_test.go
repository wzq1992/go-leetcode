package homework

import "testing"

func TestConstructor(t *testing.T) {
	circularDeque := Constructor(10)
	t.Log(circularDeque)
}

func TestMyCircularDeque_InsertFront(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	for i := 0; i < k; i++ {
		if got := circularDeque.InsertFront(i); got != true {
			t.Error("insert front error")
		}
	}

	if !circularDeque.IsFull() {
		t.Error("expect deque is full, but not")
	}

	if got := circularDeque.InsertFront(k); got != false {
		t.Error("deque is full, but insert success")
	}
}

func TestMyCircularDeque_InsertLast(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	for i := 0; i < k; i++ {
		if got := circularDeque.InsertLast(i); got != true {
			t.Error("insert last error")
		}
	}

	if !circularDeque.IsFull() {
		t.Error("expect deque is full, but not")
	}

	if got := circularDeque.InsertLast(k); got != false {
		t.Error("deque is full, but insert success")
	}
}

func TestMyCircularDeque_DeleteFront(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	for i := 0; i < k; i++ {
		if got := circularDeque.InsertLast(i); got != true {
			t.Error("insert error")
		}
	}

	if !circularDeque.IsFull() {
		t.Error("expect deque is full, but not")
	}

	for i := 0; i < k; i++ {
		if got := circularDeque.DeleteFront(); got != true {
			t.Error("delete front error")
		}
	}

	if !circularDeque.IsEmpty() {
		t.Error("expect deque is empty, but not")
	}

	if got := circularDeque.DeleteFront(); got != false {
		t.Error("deque is empty, but delete success")
	}
}

func TestMyCircularDeque_DeleteLast(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	for i := 0; i < k; i++ {
		if got := circularDeque.InsertLast(i); got != true {
			t.Error("insert error")
		}
	}

	if !circularDeque.IsFull() {
		t.Error("expect deque is full, but not")
	}

	for i := 0; i < k; i++ {
		if got := circularDeque.DeleteLast(); got != true {
			t.Error("delete last error")
		}
	}

	if !circularDeque.IsEmpty() {
		t.Error("expect deque is empty, but not")
	}

	if got := circularDeque.DeleteLast(); got != false {
		t.Error("deque is empty, but delete success")
	}
}

func TestMyCircularDeque_GetFront(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	if got := circularDeque.GetFront(); got != -1 {
		t.Error("empty deque get front must return -1")
	}

	for i := 0; i < k; i++ {
		if got := circularDeque.InsertLast(i + 1); got != true {
			t.Error("insert error")
		}
	}

	front := circularDeque.GetFront()
	if front != 1 {
		t.Errorf("expect front is 1, but not")
	}
}

func TestMyCircularDeque_GetRear(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	if got := circularDeque.GetRear(); got != -1 {
		t.Error("empty deque get rear must return -1")
	}

	for i := 0; i < k; i++ {
		if got := circularDeque.InsertLast(i + 1); got != true {
			t.Error("insert error")
		}
	}

	rear := circularDeque.GetRear()
	if rear != k {
		t.Errorf("expect rear is %d, but not", k)
	}
}

func TestMyCircularDeque_IsEmpty(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	if !circularDeque.IsEmpty() {
		t.Error("expect deque is empty, but not")
	}
}

func TestMyCircularDeque_IsFull(t *testing.T) {
	k := 10
	circularDeque := Constructor(k)

	for i := 0; i < k; i++ {
		if got := circularDeque.InsertFront(i + 1); got != true {
			t.Error("insert error")
		}
	}

	if !circularDeque.IsFull() {
		if !circularDeque.IsEmpty() {
			t.Error("expect deque is full, but not")
		}
	}
}
