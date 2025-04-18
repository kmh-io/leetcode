package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOneIII(t *testing.T) {
	testcases := []struct {
		name string
		num  []int
		k    int
		want int
	}{
		{
			name: "case 1",
			num:  []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:    2,
			want: 6,
		},
		{
			name: "case 2",
			num:  []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:    3,
			want: 10,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := longestOnes(tc.num, tc.k)

			assert.Equal(t, tc.want, actual)
		})
	}
}
