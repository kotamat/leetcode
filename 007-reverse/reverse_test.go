package _07_reverse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	x := 123
	answer := 321
	assert.Equal(t, answer, reverse(x))
}

func TestMinus(t *testing.T) {
	x := -123
	answer := -321
	assert.Equal(t, answer, reverse(x))
}

func TestLarge(t *testing.T) {
	x := 1534236469
	answer := 0
	assert.Equal(t, answer, reverse(x))
}

func BenchmarkCode(t *testing.B) {
	x := 123
	for i := 0; i < t.N; i++ {
		reverse(x)
	}
}