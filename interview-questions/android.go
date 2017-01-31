package main

import "fmt"

func getJumps() [][]int {
	x := make([][]int, 10)
	for i := range x {
		x[i] = make([]int, 10)
	}
	x[1][3], x[3][1] = 2, 2
	x[1][7], x[7][1] = 4, 4
	x[1][9], x[9][1], x[8][2], x[2][8] = 5, 5, 5, 5
	x[3][7], x[7][3], x[4][6], x[6][4] = 5, 5, 5, 5
	x[3][9], x[9][3] = 6, 6
	x[7][9], x[9][7] = 8, 8
	return x
}

func search(pattern map[int]bool, jumps [][]int, last, m, n int, total *int) {
	if len(pattern) > n {
		return
	}
	if len(pattern) >= m {
		(*total) += 1
	}
	for next := 1; next <= 9; next++ {
		req := jumps[last][next]
		if pattern[next] || (req > 0 && !pattern[req]) { // already used
			continue
		}
		pattern[next] = true
		search(pattern, jumps, next, m, n, total)
		delete(pattern, next)
	}
}

func numberOfPatterns(m int, n int) int {
	jumps := getJumps()
	total := 0
	pattern := make(map[int]bool)
	for i := 1; i <= 9; i++ {
		pattern[i] = true
		search(pattern, jumps, i, m, n, &total)
		delete(pattern, i)
	}
	return total
}

func main() {
	fmt.Println(numberOfPatterns(1, 1))
	fmt.Println(numberOfPatterns(2, 2))
	fmt.Println(numberOfPatterns(1, 2))
	fmt.Println(numberOfPatterns(1, 3))
}
