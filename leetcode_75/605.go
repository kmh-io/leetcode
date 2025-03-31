package leetcode75

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	if n == 0 {
		return true
	}

	if l == 1 {
		return flowerbed[0] == 0
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		n--
	}

	i := 1
	for i < len(flowerbed)-1 {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
		i++
	}

	if flowerbed[l-1] == 0 && flowerbed[l-2] == 0 {
		n--
	}

	return n <= 0
}
