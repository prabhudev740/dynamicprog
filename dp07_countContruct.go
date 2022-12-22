package main

import "strings"

func main() {
	// println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	// println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	// println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeeee", "eeeeeeee"}))
	println(countConstructDy("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]int{}))
	println(countConstructDy("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]int{}))
	println(countConstructDy("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeeee", "eeeeeeee"}, map[string]int{}))
}

func countConstructDy(target string, wordBank []string, memo map[string]int) int {
	if val, ok := memo[target]; ok {
		return val
	}

	if target == "" {
		return 1
	}

	var count int

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			noOfWays := countConstructDy(suffix, wordBank, memo)
			count += noOfWays
		}
	}

	memo[target] = count
	return count
}

func countConstruct(target string, wordBank []string) int {
	if target == "" {
		return 1
	}

	var count int

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			noOfWays := countConstruct(suffix, wordBank)
			count += noOfWays
		}
	}

	return count
}
