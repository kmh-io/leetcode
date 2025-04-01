package topinterview150

func maxProfitII(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	if len(prices) == 2 {
		if prices[0] >= prices[1] {
			return 0
		}
		return prices[1] - prices[0]
	}

	l, r := 0, 1
	for r < len(prices)-1 && prices[r] < prices[r+1] {
		if prices[l] >= prices[r] {
			l++
		}
		r++
	}

	if r+1 < len(prices) {
		if prices[l] >= prices[r] {
			return maxProfitII(prices[r+1:])
		}

		return prices[r] - prices[l] + maxProfitII(prices[r+1:])
	}

	if prices[l] >= prices[r] {
		return 0
	}

	return prices[r] - prices[l]
}
