package main

import (
	"fmt"
)

func main() {
	num1 := []int{1, 3}
	num2 := []int{2,4,5}

	result := findMedianSortedArrays(num1, num2)
	fmt.Println("result is ", result)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	merge := func(nums1 []int, nums2 []int) []int {
		// 空配列処理
		if len(nums1) ==0 {
			return nums2
		} else if len(nums2) == 0 {
			return nums1
		}

		//1. 2つの配列をマージして、ソートする
		min1 := nums1[0]
		max1 := nums1[len(nums1)-1]
		min2 := nums2[0]
		max2 := nums2[len(nums2)-1]

		mix := []int{}
		//	1. 2つの配列がかぶってないときは、単純に連結する(境界値含む)
		if min2 >= max1 {
			mix = append(nums1, nums2...)
		} else if min1 >= max2 {
			mix = append(nums2, nums1...)
		} else {
			//かぶっている場合は、min が少ない方の配列からforを開始し、連結していく
			base := nums1
			appended := nums2

			appendedKey := 0
			if min1 >= min2 {
				base = nums2
				appended = nums1
			}
			fmt.Println(base)
			fmt.Println(appended)

			for _, value := range base {
				for i := appendedKey; i < len(appended) ; i++ {
					if value < appended[appendedKey] {
						break
					}
					mix = append(mix, appended[appendedKey])
					appendedKey = i + 1
				}
				mix = append(mix, value)
			}

			if appendedKey < len(appended){
				mix = append(mix, appended[appendedKey:]...)
			}

			fmt.Println(mix)
		}
		return mix

	}

	mix := merge(nums1, nums2)

	//2. 1の中央値を求める

	if len(mix)%2 == 1 {
		return float64(mix[len(mix)/2])
	} else {
		return float64(mix[(len(mix)-1)/2]+mix[len(mix)/2]) / 2
	}
	return -1
}

//
