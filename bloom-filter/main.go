package main

import (
	"fmt"
	"hash/fnv"
)

// A bloom filter is a probabilistic data structure that sometimes returns false positives (FP),
// but always true negatives.
// You need generally < 10 bits per input to get a 1% FP rate.

// First, we need an m bit array and k hash functions. Typically,
// k << m. Each k hash function will map an input to a specific bit m
// with a uniform random probability.

// I used a calculator (http://hur.st/bloomfilter?n=1000000&p=0.1)
// to try to get the FP rate below 1% with 1e6 items,
// I need an m of 5,000,000.
// n = 1,000,000, p = 0.1 (1 in 10) → m = 4,792,530 (585.03KB), k = 3

const (
	m = 64 // (TODO: Support m != 64).
	k = 4
)

type hashFn func(string) uint64

type bloomFilter struct {
	seen    int
	data    uint64
	hashFns []hashFn
}

func newHashFn(n int) hashFn {
	return func(input string) uint64 {
		// The hash function just concats
		// the value of n to the input to produce
		// a "new" hash function
		in := fmt.Sprintf("%s%d", input, n)
		// Create the hash
		h := fnv.New64()
		h.Write([]byte(in))
		v := h.Sum64()
		// Set some specific bit, depending on the size of m.
		return 1 << (v % m)
	}
}

func newBloomFilter() *bloomFilter {
	var hfns []hashFn
	for i := 0; i < k; i++ {
		hfns = append(hfns, newHashFn(k))
	}
	return &bloomFilter{
		hashFns: hfns,
	}
}

func (b *bloomFilter) add(input string) {
	b.seen++
	for _, h := range b.hashFns {
		b.data |= h(input)
	}
	return
}

func (b *bloomFilter) exists(input string) bool {
	for _, h := range b.hashFns {
		if h(input)&b.data == 0 {
			return false
		}
	}
	return true
}
