package leetcode75

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func reverseVowels(s string) string {
	sbyte := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		if isVowel(sbyte[l]) && isVowel(sbyte[r]) {
			sbyte[l], sbyte[r] = sbyte[r], sbyte[l]
		}

		if !isVowel(sbyte[r]) {
			r--
		} else if !isVowel(sbyte[l]) {
			l++
		} else {
			r--
			l++
		}
	}

	return string(sbyte)
}
