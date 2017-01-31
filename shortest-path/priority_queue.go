package main

import (
	"fmt"
)

// The priority queue is shitty because it has O(n)
// pop instead of O(log(n)) like a heap-based one
// would have. TODO: Replace.
type shittyPriorityQueue map[vertex]uint64

func (s *shittyPriorityQueue) add(v vertex, weight uint64) {
	(*s)[v] = weight
	return
}

func (s *shittyPriorityQueue) isEmpty() bool {
	return len(*s) == 0
}

func (s *shittyPriorityQueue) pop() (vertex, uint64, error) {
	if s.isEmpty() {
		return "", 0, fmt.Errorf("Can't pop off an empty queue.")
	}
	// Otherwise, does an O(n) search to find the lowest
	// item, deletes it, and then returns it.
	var lowestV vertex
	var lowestW uint64 = 1<<64 - 1
	for v, w := range *s {
		if w < lowestW {
			lowestV = v
			lowestW = w
		}
	}
	delete(*s, lowestV)
	return lowestV, lowestW, nil
}
