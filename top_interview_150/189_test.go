package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)
}
