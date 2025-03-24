package leetcode75

func mergeAlternately(word1 string, word2 string) string {
	var result, i, j = make([]byte, len(word1)+len(word2)), 0, 0

	for i+j < len(result) {
		if i < len(word1) {
			result[i+j] = word1[i]
			i++
		}

		if j < len(word2) {
			result[i+j] = word2[j]
			j++
		}
	}
	return string(result)
}
