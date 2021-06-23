package homework



func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		head = l2
		l2Next := l2.Next
		l2.Next = mergeTwoLists(l1, l2Next)
	} else {
		head = l1
		l1Next := l1.Next
		l1.Next = mergeTwoLists(l2, l1Next)
	}

	return head
}
