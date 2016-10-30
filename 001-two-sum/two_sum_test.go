package _01_two_sum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4, 10}
	target := 12
	answer := []int{1, 3}
	result := twoSum(nums, target)
	assert.Equal(t, result, answer)
}

func BenchmarkTwoSum(t *testing.B) {
	nums := []int{3, 2, 4}
	target := 7
	for i := 0; i < t.N; i++ {
		twoSum(nums, target)
	}
}
