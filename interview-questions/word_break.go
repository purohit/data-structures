package main

import (
	"fmt"
)

type trie struct {
	to  map[rune]*trie
	end bool
}

func newTrie() *trie {
	return &trie{to: make(map[rune]*trie)}
}

func buildTrie(dict []string) *trie {
	top := newTrie()
	for _, w := range dict {
		cur := top
		for _, r := range w {
			if _, ok := cur.to[r]; !ok {
				cur.to[r] = newTrie()
			}
			cur = cur.to[r]
		}
		cur.end = true
	}
	return top
}

func possible(dict []string) map[rune]bool {
	p := make(map[rune]bool)
	for _, w := range dict {
		for _, r := range w {
			p[r] = true
		}
	}
	return p
}

func hasMatch(pos int, s []rune, cur, top *trie) bool {
	if _, ok := cur.to[s[pos]]; !ok {
		return false
	}
	next := cur.to[s[pos]]
	if pos == len(s)-1 {
		return next.end
	}
	matches := hasMatch(pos+1, s, next, top)
	if next.end {
		matches = matches || hasMatch(pos+1, s, top, top)
	}
	return matches
}

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	poss := possible(wordDict)
	for _, r := range s {
		if !poss[r] {
			return false
		}
	}
	t := buildTrie(wordDict)
	return hasMatch(0, []rune(s), t, t)
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
