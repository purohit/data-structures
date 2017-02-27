package main

import "fmt"

/*
Suppose you have N integers from 1 to N. We define a beautiful arrangement as an
array that is constructed by these N numbers successfully if one of the
following is true for the ith position (1 ≤ i ≤ N) in this array:

	The number at the ith position is divisible by i.
	i is divisible by the number at the ith position.

Given N, how many beautiful arrangements can you construct?
Note: N is a positive integer and will not exceed 15.
*/

func countArrangement(n int) int {
	total := 0
	fill := make([]int, n)
	avail := make(map[int]bool, n)
	for i := 1; i <= n; i++ {
		avail[i] = true
	}
	backtrack(0, &total, fill, avail)
	return total
}

// Approach: for the current open spot, try
// any of the available numbers; if it works,
// continue. Otherwise, backtrack.
func backtrack(pos int, total *int, fill []int, avail map[int]bool) {
	if pos == len(fill) {
		*total++
		return
	}
	for n, ok := range avail {
		if !ok || (n%(pos+1) != 0 && (pos+1)%n != 0) {
			continue // Not beautiful!
		}
		avail[n] = false
		fill[pos] = n
		backtrack(pos+1, total, fill, avail)
		avail[n] = true
	}
}

func main() {
	max := 15
	for i := 1; i <= max; i++ {
		fmt.Println(countArrangement(i))
	}
}
