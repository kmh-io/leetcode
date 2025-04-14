package leetcode75

func isSubsequence(s string, t string) bool {
	spointer, tpointer, count := 0, 0, len(s)

	for spointer < len(s) && tpointer < len(t) {
		for tpointer < len(t)-1 && s[spointer] != t[tpointer] {
			tpointer++
		}

		if s[spointer] == t[tpointer] {
			tpointer++
			count--
		}

		spointer++
	}

	return count == 0
}
