package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSequence(t *testing.T) {
	testcases := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "test case 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "test case 2",
			s:    "aaaaaa",
			t:    "bbaaaa",
			want: false,
		},
		{
			name: "test case 3",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSubsequence(tc.s, tc.t)

			assert.Equal(t, tc.want, actual)
		})
	}
}
