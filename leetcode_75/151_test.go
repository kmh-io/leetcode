package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "case 1",
			input: "hello world!",
			want:  "world! hello",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverseWords(tc.input)

			assert.Equal(t, tc.want, actual)
		})
	}
}
