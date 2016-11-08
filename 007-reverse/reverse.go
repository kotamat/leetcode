package _07_reverse

import (
	"math"
)

func reverse(x int) int {
	result := 0
	for ; x != 0; x /= 10 {
		result = 10 * result + x % 10
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}
	return result
}