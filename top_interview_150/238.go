package topinterview150

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = 1

	for i := 0; i < len(nums)-1; i++ {
		prefix[i+1] = nums[i] * prefix[i]
	}

	prevResult := 1
	for i := len(nums) - 1; i >= 0; i-- {
		prefix[i] *= prevResult
		prevResult *= nums[i]
	}

	return prefix
}
