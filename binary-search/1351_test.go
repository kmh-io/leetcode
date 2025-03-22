package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Count_Negative_Numbers_in_a_Sorted_Matrix(t *testing.T) {
	testcases := []struct {
		name     string
		input    [][]int
		expected int
	}{
		{
			name: "Example 1",
			input: [][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			expected: 8,
		},
		{
			name: "Example 2",
			input: [][]int{
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
				{-1, -1, -1, -1},
			},
			expected: 12,
		},
		{
			name: "Example 3",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: 0,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := countNegatives(tc.input)

			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Expected %d, got %d", tc.expected, actual))
		})
	}
}
