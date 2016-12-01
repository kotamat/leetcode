package _05_longest_palindromic_substring

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen == 1 {
		return s
	}
	LLi := 0
	LRi := 0

	calc := func(Li int, Ri int) {
		for Li >= 0 && Ri < sLen {
			if s[Li] != s[Ri] {
				break
			}
			Li--
			Ri++
		}
		if Ri - Li - 2 > LRi - LLi {
			LLi = Li + 1
			LRi = Ri - 1
		}
	}

	for key := range s {
		if sLen - (LRi - LLi) / 2 < key {
			break
		}
		// 真ん中がvalueの場合
		calc(key, key)

		// 真ん中がvalueとvalue -1の場合
		if key < 1 {
			continue
		}
		calc(key - 1, key)
	}

	return string(s[LLi:LRi + 1])
}
