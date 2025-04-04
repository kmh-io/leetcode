package leetcode75

import (
	"strings"
)

func reverseWords(s string) string {
	stringSlice := strings.Fields(s)
	l, r := 0, len(stringSlice)-1

	for l < r {
		stringSlice[l], stringSlice[r] = stringSlice[r], stringSlice[l]
		l++
		r--
	}

	return strings.Join(stringSlice, " ")
}
