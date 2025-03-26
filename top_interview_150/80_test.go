package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		expected []int
		length   int
	}{
		{
			name:     "case 1",
			input:    []int{0, 0},
			expected: []int{0, 0},
			length:   2,
		},
		{
			name:     "case 2",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
			length:   2,
		},
		{
			name:     "case 3",
			input:    []int{0, 0, 0, 1},
			expected: []int{0, 0, 1, 0},
			length:   3,
		},
		{
			name:     "case 4",
			input:    []int{0, 0, 0, 1, 1},
			expected: []int{0, 0, 1, 1, 0},
			length:   4,
		},
		{
			name:     "case 5",
			input:    []int{0, 0, 0, 1, 1, 1},
			expected: []int{0, 0, 1, 1, 0, 0},
			length:   4,
		},
		{
			name:     "case 6",
			input:    []int{1, 1, 1, 2, 2, 3},
			expected: []int{1, 1, 2, 2, 3, 0},
			length:   5,
		},
		{
			name:     "case 7",
			input:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected: []int{0, 0, 1, 1, 2, 3, 3, 0, 0},
			length:   7,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			l := removeDuplicates(tc.input)
			assert.Equal(t, tc.expected[:l], tc.input[:l])
			assert.Equal(t, tc.length, l)
		})
	}
}
