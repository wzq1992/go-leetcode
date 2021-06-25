package lessons

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var last *ListNode

	for head != nil {
		nextHead := head.Next
		head.Next, last, head = last, head, nextHead
	}

	return last
}
