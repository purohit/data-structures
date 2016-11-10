package main

type node struct {
	leaves    map[rune]*node
	wordFinal bool // a word-final marker = true means a full word is stored until here.
}

type trie struct {
	start *node
	size  int
}

func newNode() *node {
	return &node{
		leaves: make(map[rune]*node),
	}
}

// add indexes a string.
func (t *trie) add(s string) {
	curNode := t.start
	for _, c := range s {
		if _, ok := curNode.leaves[c]; !ok {
			curNode.leaves[c] = newNode()
		}
		curNode = curNode.leaves[c]
	}
	curNode.wordFinal = true // Add end marker so we know it's a word.
	t.size++
}

// search returns if s was found, as well as how many operations
// it took to find or fail fast.
func (t *trie) search(s string) (found bool, lookups int) {
	curNode := t.start
	for _, c := range s {
		lookups++
		if _, ok := curNode.leaves[c]; !ok {
			return
		}
		curNode = curNode.leaves[c]
	}
	lookups++
	found = curNode.wordFinal // At the end, make sure there's an end-marker.
	return
}

func newTrie() *trie {
	return &trie{
		start: newNode(),
	}
}
