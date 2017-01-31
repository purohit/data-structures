package main

import (
	//"fmt"
	"math"
	"strconv"
)

func isPalindrome(n int) bool {
	for s, i := strconv.Itoa(n), 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func largestPalindrome(n int) int {
	bestProd := 0
	start := int(math.Pow(10, float64(n))) - 1
	stop := int(math.Pow(10, float64(n-1)))
	for a := start; a >= stop; a-- {
		for b := a; b >= stop; b-- {
			if p := a * b; p > bestProd && isPalindrome(p) {
				bestProd = p
			}
		}
	}
	return bestProd % 1337
}

func largestPalindrome2(n int) int {
	bestProd := 0
	start := int(math.Pow(10, float64(n))) - 1
	stop := int(math.Pow(10, float64(n-1)))
	for a := stop; a <= start; a++ {
		for b := stop; b <= start; b++ {
			if p := a * b; p > bestProd && isPalindrome(p) {
				bestProd = p
			}
		}
	}
	return bestProd % 1337
}

func main() {
	// 9, 987, 123
}
