package _2

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	input := [][]int{
		[]int{0,0},
		[]int{1,0},
		[]int{2,0},
	}
	output := 2
	assert.Equal(t, output, numberOfBoomerangs(input))
}
