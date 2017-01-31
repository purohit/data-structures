package main

import (
	"fmt"
)

type vertex string

// The integer here is the distance to the given vertex
type edges map[vertex]uint64

type graph struct {
	vertices map[vertex]edges
}

func newGraph() *graph {
	return &graph{
		vertices: make(map[vertex]edges),
	}
}

type path struct {
	vertices []vertex
	distance uint64
}

func (g *graph) shortestFrom(src, dest vertex) (*path, error) {
	p := &path{}
	var best = make(edges)
	var q = make(shittyPriorityQueue)
	// First, let's check to make sure the src and dest are in the graph.
	for _, x := range []vertex{src, dest} {
		if _, ok := g.vertices[x]; !ok {
			return nil, fmt.Errorf("%s is not in graph", x)
		}
	}
	// Special case:
	if src == dest {
		return p, nil
	}
	// Now, for each one of the neighbors of the current source,
	// 1) If we've seen it before, skip it
	// 2) If we haven't seen it before, update the
	// let's update the new weight as long as it's shorter
	// than the current distance to the graph.
	q.add(src, 0)
	best[src] = 0
	for !q.isEmpty() {
		v, d, _ := q.pop()
		fmt.Printf("\nPopped %s with dist %d", v, d)
		if v == dest {
			p.distance = d
			return p, nil
		}
		for neighbor, distTo := range g.vertices[v] {
			// Already saw this neighbor, can't find
			// a shorter path because of the min queue.
			cur := d + distTo
			old, seen := best[neighbor]
			if !seen {
				old = 1<<64 - 1
			}
			if cur < old {
				best[neighbor] = cur
			}
			if !seen {
				q.add(neighbor, cur)
				fmt.Printf("\nEnqueued %s with dist %d", neighbor, cur)
			}
		}
	}
	return p, nil
}

func main() {
	// We're going to need a priority queue, as well.

}
