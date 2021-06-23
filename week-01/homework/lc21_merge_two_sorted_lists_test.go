package homework

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			"Test case 1",
			CreateLinkedList([]int{1, 2, 4}),
			CreateLinkedList([]int{1, 3, 4}),
			CreateLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			"Test case 2",
			CreateLinkedList([]int{}),
			CreateLinkedList([]int{}),
			CreateLinkedList([]int{}),
		},

		{
			"Test case 3",
			CreateLinkedList([]int{2, 3, 5, 8}),
			CreateLinkedList([]int{1, 4, 5, 9}),
			CreateLinkedList([]int{1, 2, 3, 4, 5, 5, 8, 9}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.l1, tt.l2); !reflect.DeepEqual(
				GetLinkedListVals(got), GetLinkedListVals(tt.expected)) {

				t.Errorf("mergeTwoLists(%v, %v) = %v, want %v",
					GetLinkedListVals(tt.l1), GetLinkedListVals(tt.l2),
					GetLinkedListVals(got), tt.expected)
			}
		})
	}
}
