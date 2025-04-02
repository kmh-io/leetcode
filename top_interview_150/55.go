package topinterview150

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func canJump(nums []int) bool {
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
	}
	return true
}
