package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	testcases := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			"case 1",
			[]string{"This", "is", "an", "example", "of", "text", "justification."},
			16,
			[]string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			"case 2",
			[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
			16,
			[]string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			"case 3",
			[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			16,
			[]string{"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fullJustify(tc.want, tc.maxWidth)

			assert.Equal(t, tc.want, actual)
		})
	}
}
