package leetcode75

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxIndex := 0
	for i, c := range candies {
		if c > candies[maxIndex] {
			maxIndex = i
		}
	}

	for i, c := range candies {
		if c+extraCandies >= candies[maxIndex] {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}
