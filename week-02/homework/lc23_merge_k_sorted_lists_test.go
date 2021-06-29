package homework

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		in       []*ListNode
		expected *ListNode
	}{
		{
			"Test case 1",
			[]*ListNode{CreateLinkedList([]int{1, 2, 4}), CreateLinkedList([]int{1, 3, 4})},
			CreateLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			"Test case 2",
			[]*ListNode{},
			nil,
		},
		{
			"Test case 2",
			[]*ListNode{nil, nil},
			nil,
		},
		{
			"Test case 3",
			[]*ListNode{
				CreateLinkedList([]int{2, 3, 5, 8}),
				CreateLinkedList([]int{1, 4, 5, 9}),
				CreateLinkedList([]int{5, 6, 10, 11}),
			},
			CreateLinkedList([]int{1, 2, 3, 4, 5, 5, 5, 6, 8, 9, 10, 11}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.in); !reflect.DeepEqual(
				GetLinkedListVals(got), GetLinkedListVals(tt.expected)) {

				str := "mergeKLists("
				for _, in := range tt.in {
					str += fmt.Sprintf("%v, ", GetLinkedListVals(in))
				}

				str += ") = "
				str += fmt.Sprintf("%v, want %v", GetLinkedListVals(got), GetLinkedListVals(tt.expected))

				t.Error(str)
			}
		})
	}
}
