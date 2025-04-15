package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHindex(t *testing.T) {
	testcases := []struct {
		name   string
		inputs []int
		want   int
	}{
		{
			"case 1",
			[]int{3, 0, 6, 1, 5},
			3,
		},
		{
			"case 2",
			[]int{1, 3, 1},
			1,
		},
		{
			"case 3",
			[]int{0, 0, 2},
			1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expected := hIndex(tc.inputs)
			assert.Equal(t, tc.want, expected)
		})
	}
}
