package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "case 1",
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "case 2",
			input:    []int{2, 1, 4},
			expected: 3,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxProfit(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
