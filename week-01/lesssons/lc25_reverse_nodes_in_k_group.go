package lesssons

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	protect := &ListNode{0, head}

	// 分组(找到每一组的开始、结尾), 按组遍历
	last := protect

	for head != nil {
		end := getEnd(head, k)
		if end == nil {
			break
		}

		nextGroupHead := end.Next
		reverseListGroup(head, end)

		last.Next = end
		head.Next = nextGroupHead

		last = head
		head = nextGroupHead
	}

	return protect.Next
}

func getEnd(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			break
		}
		head = head.Next
	}

	return head
}

func reverseListGroup(head, end *ListNode) {
	if head == end {
		return
	}

	last := head
	head = head.Next

	for head != end {
		nextHead := head.Next
		head.Next, last, head = last, head, nextHead
	}

	end.Next = last
}
