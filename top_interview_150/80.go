package topinterview150

func removeDuplicates(nums []int) int {
	left, right, result, fullIndex := 0, 1, 1, 0

	for right < len(nums) {

		if nums[left] == nums[right] {
			if fullIndex == 0 {
				fullIndex = right
				left++
				nums[left] = nums[right]
				result++
			}
		} else {
			if nums[fullIndex] != nums[right] {
				left++
				nums[left] = nums[right]
				fullIndex = 0
				result++
			}
		}

		right++
	}

	return result
}
