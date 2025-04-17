package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandy(t *testing.T) {
	testcases := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"case 1",
			[]int{1, 2, 87, 87, 87, 2, 1},
			13,
		},
		{
			"case 1",
			[]int{1, 3, 4, 5, 2},
			11,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := candy(tc.input)

			assert.Equal(t, tc.want, actual)
		})
	}
}
