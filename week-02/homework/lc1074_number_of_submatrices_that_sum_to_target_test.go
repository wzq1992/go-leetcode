package homework

import "testing"

func Test_numSubmatrixSumTarget(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]int
		target   int
		expected int
	}{
		{"nil slice", nil, 0, 0},
		{"one value slice", [][]int{{0}}, 0, 1},
		{"more value slice", [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 2, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := numSubmatrixSumTarget(tc.matrix, tc.target); got != tc.expected {
				t.Errorf("numSubmatrixSumTarget(%v, %v) = %v, want = %v", tc.matrix, tc.target, got, tc.expected)
			}
		})
	}
}
