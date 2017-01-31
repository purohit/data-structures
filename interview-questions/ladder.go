package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		words[w] = true
	}
	if beginWord == endWord || !words[endWord] {
		return 0
	}
	// All words have the same length! Do a BFS by performing manipulations one at a time.
	// If the word
	type wd struct {
		dist int
		word string
	}
	queue := []wd{wd{dist: 1, word: beginWord}}
	var cur wd
	for len(queue) > 0 {
		cur, queue = queue[0], queue[1:]
		if cur.word == endWord {
			return cur.dist
		}
		// Try machinations thereof we haven't seen yet
		for i := 0; i < len(cur.word); i++ {
			v := []byte(cur.word)
			for d := 0; d < 26; d++ {
				v[i] = 'a' + byte(d)
				variation := string(v)
				if !words[variation] {
					continue
				}
				queue = append(queue, wd{dist: cur.dist + 1, word: variation})
				delete(words, variation)
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog", "fog", "xxy"}))
	fmt.Println(ladderLength("hit", "nog", []string{"hot", "dot", "dog", "lot", "log", "cog", "fog", "xxy"}))
	fmt.Println(ladderLength("hit", "xxy", []string{"hot", "dot", "dog", "lot", "log", "cog", "fog", "xxy"}))
}
