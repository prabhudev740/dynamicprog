package main

import "fmt"

func main() {
	// fmt.Println(bestSum(7, []int{5, 4, 7}))
	// fmt.Println(bestSum(8, []int{2, 3, 5}))
	// fmt.Println(bestSum(8, []int{1, 4, 5}))
	// fmt.Println(bestSum(8, []int{5}))
	// fmt.Println(bestSum(100, []int{1, 2, 5, 25}))

	fmt.Println(bestSumDy(7, []int{5, 4, 7}, map[int][]int{}))
	fmt.Println(bestSumDy(8, []int{2, 3, 5}, map[int][]int{}))
	fmt.Println(bestSumDy(8, []int{1, 4, 5}, map[int][]int{}))
	fmt.Println(bestSumDy(8, []int{5}, map[int][]int{}))
	fmt.Println(bestSumDy(100, []int{1, 2, 5, 25}, map[int][]int{}))
}

func bestSum(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	// shortest := new([]int)
	var shortest []int

	for _, val := range nums {
		rem := target - val

		remRes := bestSum(rem, nums)
		if remRes != nil {
			cobination := append(remRes, val)
			if shortest == nil || len(cobination) < len(shortest) {
				shortest = cobination
			}
		}
	}

	return shortest
}

func bestSumDy(target int, nums []int, memo map[int][]int) []int {
	if val, ok := memo[target]; ok {
		return val
	}

	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	// shortest := new([]int)
	var shortest []int

	for _, val := range nums {
		rem := target - val

		remRes := bestSumDy(rem, nums, memo)
		if remRes != nil {
			cobination := append(remRes, val)
			if shortest == nil || len(cobination) < len(shortest) {
				shortest = cobination
			}
		}
	}

	memo[target] = shortest
	return shortest
}
