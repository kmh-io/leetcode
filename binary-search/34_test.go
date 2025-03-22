package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Find_First_and_Last_Position_of_Element_in_Sorted_Array(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "single target",
			input:    []int{1},
			target:   1,
			expected: []int{0, 0},
		},
		{
			name:     "target not found",
			input:    []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "target found",
			input:    []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "target found at end",
			input:    []int{1, 4},
			target:   4,
			expected: []int{1, 1},
		},
		{
			name:     "target found at end",
			input:    []int{1, 2, 3},
			target:   3,
			expected: []int{2, 2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := searchRange(tc.input, tc.target)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
