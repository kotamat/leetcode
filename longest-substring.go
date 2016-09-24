package main

import (
	"fmt"
)

func main() {
	s := "ddrewjifodabre"
	lenght := lengthOfLongestSubstring(s)
	fmt.Println(lenght)
}

func lengthOfLongestSubstring(s string) int {
	current := []rune{}
	longest := current

	strpos := func(target rune, source []rune) int {
		for key, value := range source {
			if value == target {
				return key
			}
		}
		return -1
	}

	for i, r := range s {
		if index := strpos(r, current); index >= 0 {
			if len(current) > len(longest) {
				longest = current
			}
			current = current[index+1:]
			//fmt.Println(string(r), " match index is ", index)
		}
		current = append(current, r)
		//fmt.Println("current is ", string(current))
		//fmt.Println("longest is ", string(longest))

		//fmt.Println(i, r)
	}
	if len(current) > len(longest) {
		longest = current
	}

	return len(longest)

}
