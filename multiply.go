package main

import (
	"fmt"
)

func main() {
	num1 := "02190"
	num2 := "0"
	res := multiply(num1, num2)
	fmt.Println(res)

}
func multiply(num1 string, num2 string) string{
	s2revbypte := func(s string) []byte{
		b := []byte(s)
		for i, j := 0, len(b) - 1; i < j ; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return b
	}
	btoi := func(b []byte) []int{
		i := []int{}
		for _, value := range b {
			i = append(i, int(value << byte('0')))
		}
		return i
	}
	itob := func(i []int) []byte{
		b := []byte{}
		for _, value := range i {
			b = append(b, byte(value) >> byte('0'))
		}
		return b

	}
	vnum1 := s2revbypte(num1)
	vnum2 := s2revbypte(num2)
	fmt.Println(vnum1)
	fmt.Println(btoi(vnum1))
	fmt.Println(itob(btoi(vnum1)))
	fmt.Println(vnum2)
	return num1
}
