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

func (t *trie) add(s string) {
	curNode := t.start
	for _, c := range s {
		if _, ok := curNode.leaves[c]; !ok {
			curNode.leaves[c] = newNode()
		}
		curNode = curNode.leaves[c]
	}
	// Add end marker so we know it's a word.
	// Let's pick something special, as long as there aren't
	// conflicts with the dictionary, we don't have to worry.
	curNode.wordFinal = true
	t.size++
}

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
	// At the end, make sure there's an end-marker.
	found = curNode.wordFinal
	return
}

func newTrie() *trie {
	return &trie{
		start: newNode(),
	}
}
