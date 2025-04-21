package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubArray(t *testing.T) {
	testcases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "case 1",
			input: []int{1, 1, 0, 1},
			want:  3,
		},
		{
			name:  "case 2",
			input: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			want:  5,
		},
		{
			name:  "case 3",
			input: []int{1, 0, 1, 1, 1, 0, 1, 1, 0, 1},
			want:  5,
		},
		{
			name:  "case 4",
			input: []int{0, 0, 0},
			want:  0,
		},
		{
			name:  "case 5",
			input: []int{1, 0, 0, 0, 0},
			want:  1,
		},
		{
			name:  "case 6",
			input: []int{1, 1, 0, 0, 1, 1, 1, 0, 1},
			want:  4,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := longestSubarray(tc.input)

			assert.Equal(t, tc.want, actual)
		})
	}
}
