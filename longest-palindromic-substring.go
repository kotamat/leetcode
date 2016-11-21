package main

import (
	"fmt"
)

func main() {
	min1 := 3
	min2_10 := 1
	min11 := 2
	S := 20

	restCent := S
	i := 0
	for ; restCent > 0; i++ {
		switch {
		case i <= 0:
			restCent -= min1
		case i >= 1 && i <= 9:
			restCent -= min2_10
		default:
			restCent -= min11
		}
	}
	result := i

	fmt.Println(result)

}
