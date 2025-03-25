package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	testcases := []struct {
		name     string
		num1     []int
		m        int
		num2     []int
		n        int
		expected []int
	}{
		{
			name:     "case 1",
			num1:     []int{1, 2, 3, 0, 0, 0},
			m:        3,
			num2:     []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, tc := range testcases {
		merge(tc.num1, tc.m, tc.num2, tc.n)
		assert.Equal(t, tc.expected, tc.num1)
	}
}
