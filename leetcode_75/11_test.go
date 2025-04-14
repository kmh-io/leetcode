package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	testcases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test case 1",
			input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:  49,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxArea(tc.input)

			assert.Equal(t, tc.want, actual)
		})
	}
}
