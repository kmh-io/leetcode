package leetcode75

// 0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1
func longestOnes(nums []int, k int) int {
	result, l, r, flip := 0, 0, 0, 0

	for r < len(nums) {
		if nums[r] == 1 {
			result++
		} else {
			if flip < k {
				result++
				flip++
			} else {
				result = max(result, longestOnes(nums[l+1:], k))
				break
			}
		}
		r++
	}

	return result
}
