package _03_longest_substring

func lengthOfLongestSubstring(s string) int {
	current := []rune{}
	longest := current

	strpos := func(target rune, source []rune) int {
		for key, value := range source {
			if value == target {
				return key
			}
		}
		return -1
	}

	for _, r := range s {
		if index := strpos(r, current); index >= 0 {
			if len(current) > len(longest) {
				longest = current
			}
			current = current[index + 1:]
		}
		current = append(current, r)
	}
	if len(current) > len(longest) {
		longest = current
	}

	return len(longest)

}
