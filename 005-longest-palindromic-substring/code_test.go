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
