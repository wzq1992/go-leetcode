package homework

func NewLinkedListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func CreateLinkedList(vals []int) *ListNode {
	var head *ListNode

	for i := len(vals) - 1; i >= 0; i-- {
		node := NewLinkedListNode(vals[i])
		node.Next = head
		head = node
	}

	return head
}

func GetLinkedListVals(head *ListNode) []int {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	return vals
}
