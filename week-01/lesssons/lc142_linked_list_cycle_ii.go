package lesssons

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}

			return head
		}
	}

	return nil
}
