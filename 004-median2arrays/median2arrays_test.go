package _04_median2arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	num1 := []int{1, 3}
	num2 := []int{2,4,5}

	result := findMedianSortedArrays(num1, num2)
	answer := 3.0
	assert.Equal(t, result, answer)
}

func BenchmarkCode(t *testing.B)  {
	num1 := []int{1, 3}
	num2 := []int{2,4,5}
	for i := 0; i < t.N; i++ {
		findMedianSortedArrays(num1, num2)
	}
}
