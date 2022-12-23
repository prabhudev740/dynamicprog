package main

import (
	"strings"
)

func main() {
	println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	println(allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	println(allConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeeee", "eeeeeeee"}))
	// println(allConstructDy("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]int{}))
	// println(allConstructDy("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]int{}))
	// println(allConstructDy("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeeee", "eeeeeeee"}, map[string]int{}))
}

func allConstruct(target string, wordBank []string) [][]string {
	if target == "" {
		return [][]string{[]string{}}
	}

	var result [][]string

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			suffix := target[len(word):]
			suffixWays := allConstruct(suffix, wordBank)
			for _, ways := range suffixWays {
				result = append(result, append([]string{word}, ways...))
			}
		}
	}
	return result
}
