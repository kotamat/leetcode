package _08_atoi

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	s := "  0342943856"
	i := 342943856
	assert.Equal(t, i, myAtoi(s))
}
func TestCode2(t *testing.T) {
	s := "  -0012a42"
	i := -12
	assert.Equal(t, i, myAtoi(s))
}
func TestCode3(t *testing.T) {
	s := "+-12"
	i := 0
	assert.Equal(t, i, myAtoi(s))
}
func TestCode4(t *testing.T) {
	s := "2147483648"
	i := 2147483647
	assert.Equal(t, i, myAtoi(s))
}

func BenchmarkCode(t *testing.B) {
	s := "2147483648"
	for i := 0; i< t.N; i++ {
		myAtoi(s)
	}
}
