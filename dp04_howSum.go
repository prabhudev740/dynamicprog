package main

import (
	"fmt"
)

func main() {
	fmt.Println(howSum(7, []int{2, 3}))
	fmt.Println(howSum(7, []int{5, 3, 4, 7}))
	fmt.Println(howSum(7, []int{2, 4}))
	fmt.Println(howSum(8, []int{2, 3, 5}))
	fmt.Println(howSum(0, []int{2, 3, 5}))
	// // fmt.Println(howSum(320, []int{7, 14}))

	fmt.Println(howSumDy(7, []int{2, 3}, map[int][]int{}))
	fmt.Println(howSumDy(7, []int{5, 3, 4, 7}, map[int][]int{}))
	fmt.Println(howSumDy(7, []int{2, 4}, map[int][]int{}))
	fmt.Println(howSumDy(8, []int{2, 3, 5}, map[int][]int{}))
	fmt.Println(howSumDy(0, []int{2, 3, 5}, map[int][]int{}))
	fmt.Println(howSumDy(320, []int{7, 14}, map[int][]int{}))
}

func howSum(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, val := range nums {
		rem := target - val
		resRem := howSum(rem, nums)

		if resRem != nil {
			return append(resRem, val)
		}
	}

	return nil
}

func howSumDy(target int, nums []int, memo map[int][]int) []int {
	if val, ok := memo[target]; ok {
		return val
	}

	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	for _, val := range nums {
		rem := target - val
		remRes := howSumDy(rem, nums, memo)

		if remRes != nil {
			memo[target] = append(remRes, val)
			return memo[target]
		}
	}

	memo[target] = nil
	return nil
}
