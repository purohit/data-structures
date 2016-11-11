package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := newBloomFilter()

	in := []string{
		"a", "b", "c", "d", "e", "f", "g",
	}

	for _, n := range in {
		bf.add(n)
	}

	for _, n := range in {
		if found := bf.exists(n); !found {
			t.Errorf("Should have found %s", n)
		}
	}

	n := 10
	fps := 0
	fpMaxRate := 0.1

	for i := 0; i < n; i++ {
		gen := fmt.Sprintf("%d", rand.Uint32())
		if found := bf.exists(gen); found {
			fps++
		}
	}

	if fpr := float64(fps) / float64(n); fpr > fpMaxRate {
		t.Errorf("False positive rate too high: %.2f", fpr)
	}
}
