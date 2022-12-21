package main

import "fmt"

func main() {
	// fmt.Println(canSum(7, []int{2, 3}))
	// fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	// fmt.Println(canSum(7, []int{2, 4}))
	// fmt.Println(canSum(8, []int{2, 3, 5}))
	// fmt.Println(canSum(320, []int{7, 14}))

	fmt.Println(canSumDy(7, []int{2, 3}, map[int]bool{}))
	fmt.Println(canSumDy(7, []int{5, 3, 4, 7}, map[int]bool{}))
	fmt.Println(canSumDy(7, []int{2, 4}, map[int]bool{}))
	fmt.Println(canSumDy(8, []int{2, 3, 5}, map[int]bool{}))
	fmt.Println(canSumDy(320, []int{7, 14}, map[int]bool{}))
}

func canSum(targetSum int, nums []int) bool {
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, val := range nums {
		reminder := targetSum - val
		if canSum(reminder, nums) {
			return true
		}
	}

	return false
}

func canSumDy(targetSum int, nums []int, memo map[int]bool) bool {
	if val, ok := memo[targetSum]; ok {
		return val
	}

	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, val := range nums {
		reminder := targetSum - val
		if canSumDy(reminder, nums, memo) {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return false
}
