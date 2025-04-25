package leetcode75

import (
	"slices"
	"sort"
)

// abc and bca
func closeStrings(word1 string, word2 string) bool {
	freq1, freq2, a := make([]int, 26), make([]int, 26), rune(97)

	for _, c := range word1 {
		freq1[c-a]++
	}

	for _, c := range word2 {
		freq2[c-a]++
	}

	for i, v := range freq1 {
		if (v > 0 && freq2[i] == 0) || (freq2[i] > 0 && v == 0) {
			return false
		}
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	return slices.Equal(freq1, freq2)
}
