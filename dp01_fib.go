package main

import "fmt"

func main() {
	// fmt.Println(fib(5))
	// fmt.Println(fib(7))
	// fmt.Println(fib(10))
	memo := map[int]int{}
	fmt.Println(fibDy(3, memo))
	fmt.Println(fibDy(4, memo))
	fmt.Println(fibDy(10, memo))
	fmt.Println(fibDy(50, memo))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibDy(n int, memo map[int]int) int {
	if n <= 2 {
		return 1
	}
	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibDy(n-1, memo) + fibDy(n-2, memo)

	return memo[n]
}

// 					5
// 			4				3
// 		3		2		2		1
// 2		1
