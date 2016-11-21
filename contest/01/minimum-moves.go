package _1

import (
	"math"
)

func minMoves(nums []int) int {
	for {
		max := math.MinInt8
		index := 0
		finish := true
		for i, value := range nums {
			if max < value {
				max = value
				index = i
			}
			if max != math.MinInt8 && max != value {
				finish = false
			}
		}
		if finish {
			return max
		}
		for key, _:= range nums {
			if key != index {
				nums[key] += 1
			}
		}
	}
	return 0
}
