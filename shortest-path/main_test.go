package main

import (
	"testing"
)

func TestShortestPath(t *testing.T) {
	g := newGraph()
	g.vertices = map[vertex]edges{
		"A": edges{"B": 1, "F": 3},
		"B": edges{"A": 1, "C": 3},
		"C": edges{"B": 3, "D": 2, "F": 16},
		"D": edges{"C": 2, "E": 1},
		"E": edges{"D": 1, "F": 1},
		"F": edges{"E": 1, "A": 3, "C": 16},
	}
	p, err := g.shortestFrom("A", "D")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", *p)
}
