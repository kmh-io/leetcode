package binarysearch

import "sort"

type Interval struct {
	interval []int
	index    int
}

type Intervals []Interval

func (in Intervals) Less(i, j int) bool {
	return in[i].interval[0] <= in[j].interval[0]
}

func (in Intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in Intervals) Len() int {
	return len(in)
}

func findRightInterval(intervals [][]int) []int {
	intervalList := make(Intervals, len(intervals))

	for i, v := range intervals {
		intervalList[i] = Interval{interval: v, index: i}
	}

	sort.Sort(intervalList)
	result := make([]int, len(intervalList))
	for i := range len(intervalList) {
		index := sort.Search(len(intervalList), func(j int) bool {
			return intervalList[j].interval[0] >= intervalList[i].interval[1]
		})

		if index == len(intervalList) {
			result[intervalList[i].index] = -1
		} else {
			result[intervalList[i].index] = intervalList[index].index
		}
	}

	return result
}
