package leetcode75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	testcases := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "test case 1",
			input: []int{1, 2, 3, 4, 5},
			want:  true,
		},
		{
			name:  "test case 3",
			input: []int{20, 100, 10, 12, 5, 13},
			want:  true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := increasingTriplet(tc.input)

			assert.Equal(t, tc.want, actual)
		})
	}
}
