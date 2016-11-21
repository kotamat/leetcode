package _06_zigzag_conversion

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/pkg/profile"
)

func TestZigzag(t *testing.T)  {
	s := "PAYPALISHIRING"
	nRows := 6
	answer := "PINALSIGYAHRPI"
	assert.Equal(t, convert(s, nRows),answer)
}

func BenchmarkZigzag(t *testing.B)  {

	s := "PAYPALISHIRING"
	nRows := 3

	for i := 0; i < t.N; i++ {
		convert(s, nRows)
	}
}
