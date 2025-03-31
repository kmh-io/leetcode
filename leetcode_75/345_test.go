package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverVerwal(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "case 1",
			input:    "IceCreAm",
			expected: "AceCreIm",
		},
		{
			name:     "case 2",
			input:    "leetcode",
			expected: "leotcede",
		},
		{
			name:     "case 3",
			input:    "\"Slang is not suet, is it?\" Euston signals.",
			expected: "\"Slang is not suEt, is it?\" euston signals.",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverseVowels(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
