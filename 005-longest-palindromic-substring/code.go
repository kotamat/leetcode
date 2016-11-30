package _05_longest_palindromic_substring

func longestPalindrome(s string) string {
	longestPalindrome := []rune{}
	sr := []rune(s)

	for key := range sr {
		// 真ん中がvalueの場合
		palindrome := []rune{}
		Li := key
		Ri := key

		for {
			if sr[Li] == sr[Ri] {
				palindrome = sr[Li : Ri + 1]
			}
			if Li <= 0 || Ri >= len(sr) - 1 || sr[Li] != sr[Ri] {
				break
			}
			Li--
			Ri++
		}

		if len(palindrome) > len(longestPalindrome) {
			longestPalindrome = palindrome
		}

		// 真ん中がvalueとvalue -1の場合
		if key < 1 {
			continue
		}
		palindrome = []rune{}
		Li = key - 1
		Ri = key
		for {
			if sr[Li] == sr[Ri] {
				palindrome = sr[Li : Ri + 1]
			}
			if Li <= 0 || Ri >= len(sr) - 1 || sr[Li] != sr[Ri] {
				break
			}
			Li--
			Ri++
		}

		if len(palindrome) > len(longestPalindrome) {
			longestPalindrome = palindrome
		}
	}

	return string(longestPalindrome)
}
