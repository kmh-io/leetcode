package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Find_Smallest_Letter_Greater_Than_Target(t *testing.T) {
	testcases := []struct {
		name     string
		input    []byte
		target   byte
		expected byte
	}{
		{
			name:     "test case 1",
			input:    []byte{'c', 'f', 'j'},
			target:   'c',
			expected: 'f',
		},
		{
			name:     "test case 2",
			input:    []byte{'c', 'f', 'j'},
			target:   'a',
			expected: 'c',
		},
		{
			name:     "test case 3",
			input:    []byte{'x', 'x', 'y', 'y'},
			target:   'z',
			expected: 'x',
		},
		{
			name:     "test case 4",
			input:    []byte{'c', 'f', 'g'},
			target:   'd',
			expected: 'f',
		},
		{
			name:     "test case 5",
			input:    []byte{'c', 'f', 'j'},
			target:   'j',
			expected: 'c',
		},
		{
			name:     "test case 6",
			input:    []byte{'e', 'e', 'g', 'g'},
			target:   'g',
			expected: 'e',
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := nextGreatestLetter(tc.input, tc.target)
			assert.Equal(t, result, tc.expected)
		})
	}
}
