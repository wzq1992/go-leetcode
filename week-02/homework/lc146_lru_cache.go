package homework

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*Node
	head, tail *Node
}

type Node struct {
	key, value int
	prev, next *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	// 哨兵节点
	head, tail := NewNode(0, 0), NewNode(0, 0)
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		size:     0,
		cache:    map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if item, ok := l.cache[key]; !ok {
		return -1
	} else {
		// 将该节点移动队首
		l.moveToFront(item)

		return item.value
	}
}

func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.cache[key]; !ok {
		// 队列已满，使用 LRU 算法淘汰节点
		if l.isFull() {
			outNode := l.tail.prev
			delete(l.cache, outNode.key)
			l.removeNode(outNode)
			l.size--
		}

		// 不存在, 插入到队首
		node := NewNode(key, value)
		l.insertToFront(node)
		l.size++
	} else {
		// 存在, 更新节点 value 并移动到队首
		item.value = value
		l.moveToFront(item)
	}
}

func (l *LRUCache) isFull() bool {
	return l.size == l.capacity
}

func (l *LRUCache) moveToFront(node *Node) {
	l.removeNode(node)
	l.insertToFront(node)
}

func (l *LRUCache) insertToFront(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node

	l.cache[node.key] = node
}

func (l *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	delete(l.cache, node.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
