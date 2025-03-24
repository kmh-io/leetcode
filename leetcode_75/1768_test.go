package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	testcases := []struct {
		name     string
		input    [2]string
		expected string
	}{
		{
			name:     "Case 1",
			input:    [2]string{"abc", "pqr"},
			expected: "apbqcr",
		},
		{
			name:     "Case 2",
			input:    [2]string{"ab", "pqrs"},
			expected: "apbqrs",
		},
		{
			name:     "Case 2",
			input:    [2]string{"abcd", "pq"},
			expected: "apbqcd",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := mergeAlternately(tc.input[0], tc.input[1])

			assert.Equal(t, tc.expected, actual)
		})
	}
}
