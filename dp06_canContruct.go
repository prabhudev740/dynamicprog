package main

import (
	"strings"
)

func main() {
	println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeeee", "eeeeeeee"}))
	println(canConstructDy("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, map[string]bool{}))
	println(canConstructDy("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, map[string]bool{}))
	println(canConstructDy("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeeee", "eeeeeeee"}, map[string]bool{}))
}

func canConstruct(targat string, wordBank []string) bool {
	if targat == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(targat, word) == 0 {
			suffix := targat[len(word):]
			if canConstruct(suffix, wordBank) {
				return true
			}
		}
	}

	return false
}

func canConstructDy(targat string, wordBank []string, memo map[string]bool) bool {
	if val, ok := memo[targat]; ok {
		return val
	}

	if targat == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(targat, word) == 0 {
			suffix := targat[len(word):]
			if canConstructDy(suffix, wordBank, memo) {
				memo[targat] = true
				return true
			}
		}
	}

	memo[targat] = false
	return false
}
