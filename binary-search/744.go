package binarysearch

import (
	"log"
)

func searchRepeact(nums []byte, target byte, curIndex int) int {
	for curIndex+1 < len(nums) && nums[curIndex+1] == target {
		curIndex++
	}
	return curIndex
}

func search(nums []byte, target byte) (int, int, int) {
	l, r := 0, len(nums)
	for l < r {
		half := (r + l) / 2
		if nums[half] == target {
			return searchRepeact(nums, target, half), l, r
		}
		if nums[half] > target {
			r = half
		} else {
			l = half + 1
		}
	}
	return -1, l, r
}

func nextGreatestLetter(letters []byte, target byte) byte {
	index, left, right := search(letters, target)

	log.Println("left: ", left, " right: ", right)

	if index == -1 {
		if target <= letters[0] || target >= letters[len(letters)-1] {
			return letters[0]
		} else {
			return letters[right]
		}
	}

	if index == len(letters)-1 {
		return letters[0]
	}

	return letters[index+1]
}
