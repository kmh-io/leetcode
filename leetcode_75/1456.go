package leetcode75

func maxVowels(s string, k int) int {
	count := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	// leetcode
	cur := count
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			cur--
		}

		if isVowel(s[i]) {
			cur++
		}

		count = max(cur, count)
	}

	return count
}
