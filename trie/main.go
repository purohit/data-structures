package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	dict           = "/usr/share/dict/american-english"
	endMarker rune = '終'
)

type node map[rune]node

type trie struct {
	start node
	size  int
}

func (t *trie) add(s string) {
	curNode := t.start
	for _, c := range s {
		if _, ok := curNode[c]; !ok {
			curNode[c] = make(node)
		}
		curNode = curNode[c]
	}
	// Add end marker so we know it's a word.
	// Let's pick something special, as long as there aren't
	// conflicts with the dictionary, we don't have to worry.
	if _, ok := curNode[endMarker]; !ok {
		curNode[endMarker] = make(node)
	}
	t.size++
}

func (t *trie) search(s string) (found bool, lookups int) {
	curNode := t.start
	for _, c := range s {
		lookups++
		if _, ok := curNode[c]; !ok {
			return
		}
		curNode = curNode[c]
	}
	lookups++
	// At the end, make sure there's an end-marker.
	_, found = curNode[endMarker]
	return
}

func newTrie() *trie {
	return &trie{
		start: make(node),
	}
}

func main() {
	f, err := os.Open(dict)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	d := newTrie()
	// Add every word to our trie.
	for s.Scan() {
		d.add(s.Text())
	}
	log.Printf("size of dict is %d", d.size)
	// Try searching for yourself.
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		t, _ := input.ReadString('\n')
		t = strings.TrimRight(t, "\n")
		found, l := d.search(t)
		if found {
			log.Printf("Found \"%s\" after %d lookups", t, l)
		} else {
			log.Printf("Failed to find \"%s\"; bailed after %d lookups", t, l)
		}
	}
}
