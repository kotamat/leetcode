package _05_longest_palindromic_substring

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	s := "babad"
	answer := "bab"
	assert.Equal(t, answer, longestPalindrome(s))
}
func TestCode2(t *testing.T) {
	s := "cbbd"
	answer := "bb"
	assert.Equal(t, answer, longestPalindrome(s))
}
func TestCode3(t *testing.T) {
	s := "abababababa"
	answer := "abababababa"
	assert.Equal(t, answer, longestPalindrome(s))
}
func BenchmarkCode(t *testing.B) {
	for i := 0; i < t.N; i++ {
		s := "babad"
		longestPalindrome(s)
	}
}
