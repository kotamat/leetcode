package _01_two_sum

func twoSum(nums []int, target int) []int {
	pool := map[int]int{}
	for key, value := range nums {
		if i, found := pool[value] ; found {
			return []int{i, key}
		}
		pool[target - value] = key
	}
	return []int{}
}
