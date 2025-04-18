package leetcode75

func maxOperations(nums []int, k int) int {
	result, cache := 0, make(map[int]int)

	for _, num := range nums {
		cache[num]++
		sub := k - num

		if sub == num {
			if cache[num] >= 2 {
				result++
				cache[num] -= 2
			}
		} else if cache[sub] > 0 {
			result++
			cache[sub]--
			cache[num]--
		}
	}

	return result
}
