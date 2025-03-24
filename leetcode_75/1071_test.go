package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	testcases := []struct {
		name     string
		input    [2]string
		expected string
	}{
		{
			name:     "Case 1",
			input:    [2]string{"ABC", "ABC"},
			expected: "ABC",
		},
		{
			name:     "Case 2",
			input:    [2]string{"ABCABC", "ABC"},
			expected: "ABC",
		},
		{
			name:     "Case 3",
			input:    [2]string{"ABABAB", "ABAB"},
			expected: "AB",
		},
		{
			name:     "Case 4",
			input:    [2]string{"LEET", "CODE"},
			expected: "",
		},
		{
			name:     "Case 5",
			input:    [2]string{"EFGABC", "ABC"},
			expected: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := gcdOfStrings(tc.input[0], tc.input[1])

			assert.Equal(t, tc.expected, actual)
		})
	}
}
