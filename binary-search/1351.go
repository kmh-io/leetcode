package binarysearch

func countNegative(input []int) int {
	count := 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] < 0 {
			count++
		} else {
			break
		}
	}
	return count
}

func countNegatives(grid [][]int) int {
	totalCount := 0
	for _, input := range grid {
		totalCount += countNegative(input)
	}
	return totalCount
}
