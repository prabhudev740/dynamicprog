package main

import "fmt"

func main() {
	// fmt.Println(gridTraveler(1, 1))
	// fmt.Println(gridTraveler(0, 1))
	// fmt.Println(gridTraveler(1, 0))
	// fmt.Println(gridTraveler(0, 0))
	// fmt.Println(gridTraveler(2, 3))
	// fmt.Println(gridTraveler(5, 3))
	// fmt.Println(gridTraveler(4, 4))

	memo := map[string]int{}

	fmt.Println(gridTravelerDy(1, 1, memo))
	fmt.Println(gridTravelerDy(0, 1, memo))
	fmt.Println(gridTravelerDy(1, 0, memo))
	fmt.Println(gridTravelerDy(0, 0, memo))
	fmt.Println(gridTravelerDy(2, 3, memo))
	fmt.Println(gridTravelerDy(5, 3, memo))
	fmt.Println(gridTravelerDy(10, 50, memo))
}

func gridTraveler(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if n == 1 && n == 1 {
		return 1
	}

	return gridTraveler(m-1, n) + gridTraveler(m, n-1)
}

func gridTravelerDy(m, n int, memo map[string]int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if n == 1 && n == 1 {
		return 1
	}

	var key string

	if n > m {
		key = fmt.Sprint("%d %d", m, n)
	} else {
		key = fmt.Sprint("%d %d", m, n)
	}

	if val, ok := memo[key]; ok {
		return val
	}

	memo[key] = gridTravelerDy(m-1, n, memo) + gridTravelerDy(m, n-1, memo)

	return memo[key]
}
