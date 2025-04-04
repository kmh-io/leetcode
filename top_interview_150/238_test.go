package topinterview150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}
	actual := productExceptSelf(input)

	assert.Equal(t, expected, actual)
}
