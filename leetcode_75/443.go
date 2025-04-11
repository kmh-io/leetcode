package leetcode75

import (
	"fmt"
)

func compress(chars []byte) int {
	l, r, replaceIndex, length := 0, 1, 0, len(chars)

	if length == 1 {
		return length
	}

	count := 1
	for r < length {
		if chars[l] != chars[r] {

			if count != 1 {
				countstr := fmt.Sprintf("%c%d", chars[l], count)
				for _, c := range []byte(countstr) {
					chars[replaceIndex] = c
					replaceIndex++
				}
			} else {
				chars[replaceIndex] = chars[l]
				replaceIndex++
			}

			l = r
			count = 1
		} else {
			count++
		}

		r++
	}

	if count != 1 {
		countstr := fmt.Sprintf("%c%d", chars[l], count)
		for _, c := range []byte(countstr) {
			chars[replaceIndex] = c
			replaceIndex++
		}
	} else {
		chars[replaceIndex] = chars[r-1]
		return replaceIndex + 1
	}

	return replaceIndex
}
