//package _06_zigzag_conversion
package main

import (
	"github.com/pkg/profile"
)

func main() {
	profile.MemProfileRate(1)
	defer profile.Start(profile.MemProfile, profile.ProfilePath("my.prof"), profile.CPUProfile).Stop()
	s := "PAYPALISHIRING"
	nRows := 6
	convert(s, nRows)
}

func convert(s string, numRows int) string {
	l := len(s)
	if l == 0 || numRows == 1 {
		return s
	}
	sl := [][]byte{}
	result := []byte{}
	lastIndex := 0
	for l > lastIndex {
		roopMod := numRows * 2 - 2
		roopIndex := lastIndex % roopMod
		if roopIndex == 0 {
			if l - lastIndex > numRows {
				sl = append(sl, []byte(s[lastIndex:lastIndex + numRows]))
			} else {
				sl = append(sl, []byte(s[lastIndex:]))
				break
			}
			lastIndex += numRows
		} else {
			tmp := make([]byte, numRows)
			tmp[roopMod - roopIndex] = byte(s[lastIndex])
			sl = append(sl, tmp)
			lastIndex ++
		}
	}

	for j := 0; j < numRows; j++ {
		for _, str := range sl {
			if j >= len(str) {
				continue
			}
			if r := str[j]; r != 0 {
				result = append(result, str[j])
			}
		}
	}

	return string(result)
}
