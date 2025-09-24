/**
 * Merge Intervals
 */

package topinterview150

import "slices"

func mergeIntervals(intervals [][]int) [][]int {
	leftIdx, rightIdx := 0, 1

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for rightIdx < len(intervals) {
		if intervals[rightIdx][0] <= intervals[leftIdx][1] {
			intervals[leftIdx][1] = max(intervals[leftIdx][1], intervals[rightIdx][1])
		} else {
			leftIdx++
			intervals[leftIdx] = intervals[rightIdx]
		}
		rightIdx++
	}
	return intervals[:leftIdx+1]
}
