package homework

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected []int
	}{
		{"The number is 0", []int{0}, []int{1}},
		{"The number goes up one place", []int{9, 9, 9}, []int{1, 0, 0, 0}},
		{
			"The big number",
			[]int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 3},
			[]int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.in); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.in, got, tt.expected)
			}
		})
	}
}
