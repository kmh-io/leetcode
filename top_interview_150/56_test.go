package topinterview150

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	testcases := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name: "Test 1",
			intervals: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			expected: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "Test 2",
			intervals: [][]int{
				{1, 4},
				{4, 5},
			},
			expected: [][]int{
				{1, 5},
			},
		},
		{
			name: "Test 3",
			intervals: [][]int{
				{1, 4},
				{2, 3},
			},
			expected: [][]int{
				{1, 4},
			},
		},
		{
			name: "Test 4",
			intervals: [][]int{
				{1, 4},
				{0, 4},
			},
			expected: [][]int{
				{0, 4},
			},
		},
		{
			name: "Test 5",
			intervals: [][]int{
				{1, 4},
				{0, 2},
				{3, 5},
			},
			expected: [][]int{
				{0, 5},
			},
		},
		{
			name: "Test 6",
			intervals: [][]int{
				{1, 4},
				{0, 0},
			},
			expected: [][]int{
				{0, 0},
				{1, 4},
			},
		},
		{
			name: "Test 7",
			intervals: [][]int{
				{1, 4},
				{5, 6},
			},
			expected: [][]int{
				{1, 4},
				{5, 6},
			},
		},
		{
			name: "Test 8",
			intervals: [][]int{
				{2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
				{1, 10},
			},
			expected: [][]int{
				{1, 10},
			},
		},
	}

	for _, tc := range testcases {
		interval := make([][]int, len(tc.intervals))
		copy(interval, tc.intervals)
		result := mergeIntervals(tc.intervals)
		assert.Equal(t, tc.expected, result, fmt.Sprintf(
			"%s: interval: %v", tc.name, interval))
	}
}
