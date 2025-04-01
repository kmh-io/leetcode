package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitII(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"case -1",
			[]int{5, 1},
			0,
		},
		{
			"case 0",
			[]int{1, 5},
			4,
		},
		{
			"case 1",
			[]int{7, 1, 5},
			4,
		},
		{
			"case 2",
			[]int{7, 1, 5, 3},
			4,
		},
		{
			"case 3",
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			"case 4",
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			"case 5",
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			"case 6",
			[]int{4, 2, 1, 7},
			6,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maxProfitII(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
