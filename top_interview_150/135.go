package topinterview150

func candy(ratings []int) int {
	l := len(ratings)
	candies := make([]int, l, l)

	for i := range l {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		r := ratings[i-1]
		if ratings[i] > r {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := l - 2; i >= 0; i-- {
		r := ratings[i+1]
		if ratings[i] > r {
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	result := 0
	for _, candy := range candies {
		result += candy
	}

	return result
}
