package leetcode75

func increasingTriplet(nums []int) bool {
	first, second := -1, -1
	for i := 1; i < len(nums); i++ {
		if first < 0 {
			if nums[i-1] < nums[i] {
				first = i - 1
				second = i
			}
		} else {
			if nums[i] > nums[second] {
				return true
			}

			if nums[first] > nums[i] {
				first = i
			}

			if nums[first] < nums[i] && nums[i] < nums[second] {
				second = i
			}
		}
	}

	return false
}
