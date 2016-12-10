package _06_zigzag_conversion

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestZigzag(t *testing.T)  {
	s := "PAYPALISHIRING"
	nRows := 3
	answer := "PAHNAPLSIIGYIR"
	assert.Equal(t, convert(s, nRows),answer)
}
func TestZigzag2(t *testing.T)  {
	s := "Apalindromeisaword,phrase,number,orothersequenceofunitsthatcanbereadthesamewayineitherdirection,withgeneralallowancesforadjustmentstopunctuationandworddividers."

	nRows := 3
	answer := "PAHNAPLSIIGYIR"
	assert.Equal(t, convert(s, nRows),answer)
}

func BenchmarkZigzag(t *testing.B)  {

	s := "PAYPALISHIRING"
	nRows := 3

	for i := 0; i < t.N; i++ {
		convert(s, nRows)
	}
}
