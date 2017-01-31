package main

import (
	"testing"
)

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestPalindrome(3)
	}
}

func BenchmarkPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestPalindrome2(3)
	}
}
