package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	testcases := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			name:  "test case 1",
			input: []byte("abb"),
			want:  3,
		},
		{
			name:  "test case 2",
			input: []byte("aabbccc"),
			want:  6,
		},
		{
			name:  "test case 3",
			input: []byte("abbbbbbbbbbbb"),
			want:  4,
		},
		{
			name:  "test case 4",
			input: []byte("abc"),
			want:  3,
		},
		{
			name:  "test case 5",
			input: []byte("aaaaab"),
			want:  3,
		},
		{
			name:  "test case 6",
			input: []byte("aaaaba"),
			want:  4,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := compress(tc.input)

			assert.Equal(t, tc.want, actual)
		})
	}
}
