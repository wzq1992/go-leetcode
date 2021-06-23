package homework

import (
	"testing"
)

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		k        int
		expected int
	}{
		{"Test case 1", []int{1, 1, 1}, 2, 2},
		{"Test case 2", []int{}, 3, 0},
		{"Test case 3", []int{1, 2, 3, 4}, 3, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.in, tt.k); got != tt.expected {
				t.Errorf("subarraySum(%v, %v) = %v, want %v", tt.in, tt.k, got, tt.expected)
			}
		})
	}
}
