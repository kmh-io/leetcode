package topinterview150

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		lastEelm := nums[len(nums)-1]
		copy(nums[1:], nums[:])
		nums[0] = lastEelm
	}
}
