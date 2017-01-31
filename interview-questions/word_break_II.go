package main

import "fmt"
import "strings"

func breakable(s string, wordDict []string) bool {
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

func wordBreak(s string, wordDict []string) []string {
	if !breakable(s, wordDict) {
		return []string{}
	}
	words := make(map[string]bool)
	for _, w := range wordDict {
		words[w] = true
	}
	var cur, all []string
	fillResults(s, &cur, &all, words)
	return all
}

func fillResults(s string, cur, all *[]string, words map[string]bool) {
	if len(s) == 0 {
		if len(*cur) > 0 {
			(*all) = append(*all, strings.Join(*cur, " "))
		}
	}
	for i := 1; i <= len(s); i++ {
		if words[s[:i]] {
			*cur = append(*cur, s[:i])
			fillResults(s[i:], cur, all, words)
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Println(wordBreak("catsanddog", []string{"catx", "rats", "and", "sand", "dog"}))
	fmt.Println(wordBreak("", []string{"catx", "rats", "and", "sand", "dog"}))
	res := wordBreak("aaaaaaa", []string{"a", "aa", "aaaa"})
	fmt.Printf("%+v", res)
}
