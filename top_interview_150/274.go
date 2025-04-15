package topinterview150

import (
	"slices"
)

func hIndex(citations []int) int {
	if len(citations) == 1 {
		if citations[0] == 0 {
			return 0
		}
		return 1
	}

	slices.SortFunc(citations, func(a, b int) int {
		return b - a
	})

	if citations[0] == 1 {
		return 1
	}

	max := 0
	for i := len(citations); i >= 1; i-- {
		cur := citations[i-1]
		if cur >= i {
			if i > max {
				max = i
			}
		}
	}

	return max
}
