package leetcode75

import (
	"math"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	minLen := int(math.Min(float64(len(str1)), float64(len(str2))))

	for ; minLen > 0; minLen-- {
		s1, s2, base := str1, str2, str1[:minLen]
		if len(str1)%minLen == 0 && len(str2)%minLen == 0 {
			s1 = strings.ReplaceAll(s1, base, "")
			s2 = strings.ReplaceAll(s2, base, "")
		}

		if s1 == "" && s2 == "" {
			return base
		}
	}

	return ""
}
