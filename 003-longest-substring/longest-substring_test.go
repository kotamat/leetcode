package _03_longest_substring

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	answer := 3
	assert.Equal(t, answer, lengthOfLongestSubstring(s))
}

func BenchmarkLongestSubstring(t *testing.B)  {
	s := "pwwkew"
	for i := 0; i < t.N; i++ {
		lengthOfLongestSubstring(s)
	}
}
