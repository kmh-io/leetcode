package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseString(t *testing.T) {
	testcases := []struct {
		name  string
		word1 string
		word2 string
		want  bool
	}{{
		"case 1",
		"cabbba",
		"abbccc",
		true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := closeStrings(tc.word1, tc.word2)

			assert.Equal(t, tc.want, actual)
		})
	}
}
