package homework

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	testCases := []struct {
		name     string
		in       []int
		expected int
	}{
		{"empty slice nums", []int{}, 0},
		{"nil slice nums", nil, 0},
		{"more values slice nums", []int{1, 2, 1, 2, 3, 2, 5}, 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := findShortestSubArray(tc.in); got != tc.expected {
				t.Errorf("findShortestSubArray(%v) = %v, want = %v", tc.in, got, tc.expected)
			}
		})
	}
}
