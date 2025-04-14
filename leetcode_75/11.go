package leetcode75

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func maxArea(height []int) int {
	l, r, max := 0, len(height)-1, 0

	for l < r {
		minHeight := min(height[l], height[r])
		area := minHeight * (r - l)
		if area > max {
			max = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return max
}
