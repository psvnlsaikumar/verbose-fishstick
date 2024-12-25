package testing

import (
	"psvnlsaikumar/go-lc/problems/easy"
	"testing"
)

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "LC Example 1",
			nums:     []int{1, 2, 3, 4, 2, 3, 3, 5, 7},
			expected: 2,
		},
		{
			name:     "LC Example 2",
			nums:     []int{4, 5, 6, 4, 4},
			expected: 2,
		},
		{
			name:     "LC Example 3",
			nums:     []int{6, 7, 8, 9},
			expected: 0,
		},

		{
			name:     "No operations needed",
			nums:     []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "Single element array",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "All elements are the same",
			nums:     []int{5, 5, 5, 5},
			expected: 1,
		},
		{
			name:     "Multiple duplicates",
			nums:     []int{4, 3, 2, 2, 4, 1},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := easy.MinimumOperations(test.nums)
			if result != test.expected {
				t.Errorf("For input %v, expected %d but got %d", test.nums, test.expected, result)
			}
		})
	}
}
