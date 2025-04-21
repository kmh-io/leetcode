package leetcode75

import "fmt"

// 1,1,0,1
// 0,1,1,1,0,1,1,0,1
// 1,0,1,1,1,0,1,1,0,1
// 0,0,0
// 1,0,0,0,0
// 1,1,0,0,1,1,1,0,1
func longestSubarray(nums []int) int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i-1] == 1 {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i+1] == 1 {
			right[i] = right[i+1] + 1
		}
	}

	fmt.Println("left ", left, " right ", right)
	result := 0
	for i := 0; i < n; i++ {
		result = max(result, left[i]+right[i])
	}
	return result
}
