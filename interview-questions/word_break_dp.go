package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for r, i := []rune(s), 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[string(r[j:i])] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func main() {
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aa"}))
	fmt.Println(wordBreak(
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code", "fox", "loop"}))
	fmt.Println(wordBreak("leetcode", []string{""}))
	fmt.Println(wordBreak("", []string{""}))
	fmt.Println(wordBreak("", []string{"leet", "code"}))
	fmt.Println(wordBreak("leetbode", []string{"leet", "code", "fox", "loop"}))
	fmt.Println(wordBreak("latagator", []string{"leet", "code", "gator", "loop"}))
	fmt.Println(wordBreak("amanaplan", []string{"a", "man", "plan"}))
	fmt.Println(wordBreak("panama", []string{"pan", "man"}))
}
