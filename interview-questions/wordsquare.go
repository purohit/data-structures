package main

import "fmt"

type ptree map[string][]string

func prefixes(words []string) ptree {
	r := make(ptree)
	r[""] = words
	for _, w := range words {
		for i := 1; i <= len(w); i++ {
			s := w[:i]
			if _, ok := r[s]; !ok {
				r[s] = make([]string, 0)
			}
			r[s] = append(r[s], w)
		}
	}
	return r
}

func wordSquares(words []string) [][]string {
	x := len(words[0])
	if len(words) < x { // Can't build a square without >= x words!
		return nil
	}
	p := prefixes(words)
	square := make([]string, x)
	squares := make([][]string, 0)
	fillSquares(0, x, p, &square, &squares)
	return squares
}

func fillSquares(i, x int, p ptree, square *[]string, squares *[][]string) {
	if i == x { // At the end, this is a valid result
		toAdd := make([]string, len(*square))
		copy(toAdd, *square)
		*squares = append(*squares, toAdd)
		return
	}
	// Otherwise, the prefix depends on the number of words we currently have
	var prefix []byte
	for row := 0; row < i; row++ {
		prefix = append(prefix, (*square)[row][i])
	}
	validWords := p[string(prefix)]
	for _, w := range validWords {
		(*square)[i] = w
		fillSquares(i+1, x, p, square, squares)
	}
	return
}

func main() {
	in := []string{"area", "lead", "wall", "lady", "ball"}
	fmt.Println(wordSquares(in))
	in = []string{"abat", "baba", "atan", "atal"}
	fmt.Println(wordSquares(in))
}
