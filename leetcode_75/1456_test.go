package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxVowels(t *testing.T) {
	testcases := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "case 1",
			s:    "abciiidef",
			k:    3,
			want: 3,
		},
		{
			name: "case 2",
			s:    "aeiou",
			k:    2,
			want: 2,
		},
		{
			name: "case 3",
			s:    "leetcode",
			k:    3,
			want: 2,
		},
		{
			name: "case 4",
			s:    "tryhard",
			k:    4,
			want: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxVowels(tc.s, tc.k)

			assert.Equal(t, tc.want, actual)
		})
	}
}
