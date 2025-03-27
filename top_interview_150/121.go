package topinterview150

func maxProfit(prices []int) int {
	profit, l, r := 0, 0, 1

	for r < len(prices) {
		if prices[r] < prices[l] {
			l = r
		} else {
			diff := prices[r] - prices[l]
			if diff > profit {
				profit = diff
			}
		}
		r++
	}

	return profit
}
