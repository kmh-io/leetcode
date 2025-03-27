package topinterview150

func majorityElement(nums []int) int {
	n := len(nums)
	times := n / 2
	cache := make(map[int]int)
	max := 0
	for _, num := range nums {
		cache[num]++
		if cache[num] > times && cache[num] > max {
			max = num
		}
	}
	return max
}
