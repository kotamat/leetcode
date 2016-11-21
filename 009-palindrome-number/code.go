package _09_palindrome_number

import (
	"math"
)

func isPalindrome(x int) bool {
	//overflow
	if 0 > x || x > math.MaxInt32 {
		return false
	}
	l := int(math.Log10(float64(x)))
	if l < 1 {
		return true
	}
	arr := []int{}
	for i := 0; i <= l; i++ {
		if i <= l / 2 {
			arr = append(arr, x % 10)
		} else if arr[l - i] != x % 10 {
			return false
		}
		x /= 10
	}
	return true
}
