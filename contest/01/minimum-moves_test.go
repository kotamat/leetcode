package _1

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCode(t *testing.T) {
	input := []int{1,2,3}
	output := 3
	assert.Equal(t, output, minMoves(input))
}
