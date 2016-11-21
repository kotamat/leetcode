package _09_palindrome_number

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	i := -2147483648
	assert.Equal(t, false, isPalindrome(i))
}

func TestTrue(t *testing.T) {
	i := 98766789
	assert.Equal(t, true, isPalindrome(i))
}
func TestTrueOdd(t *testing.T) {
	i := 987656789
	assert.Equal(t, true, isPalindrome(i))
}
func TestZero(t *testing.T) {
	i := 0
	assert.Equal(t, true, isPalindrome(i))
}
func Test10(t *testing.T) {
	i := 10
	assert.Equal(t, false, isPalindrome(i))
}

func BenchmarkCode(t *testing.B)  {
	for i := 0; i < t.N; i++ {
		isPalindrome(9783928953)
	}
}
