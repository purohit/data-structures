package main

import "fmt"

func numIslands(grid [][]byte) int {
	var rows, cols, islands int
	rows = len(grid)
	if rows > 0 {
		cols = len(grid[0])
	}
	type pt struct{ i, j int }
	markAdjacent := func(i, j int) {
		queue := []pt{{i, j}}
		for len(queue) > 0 {
			i, j = queue[0].i, queue[0].j
			queue = queue[1:]
			grid[i][j] = 'X'
			for _, n := range []pt{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}} {
				i, j = n.i, n.j
				if i < 0 || j < 0 || i > rows-1 || j > cols-1 || grid[i][j] != '1' {
					continue
				}
				queue = append(queue, pt{i, j})
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' || grid[i][j] == 'X' {
				continue
			}
			islands++
			markAdjacent(i, j)
		}
	}
	return islands
}

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '0', '0'},
		[]byte{'0', '1', '0', '0', '0'},
		[]byte{'1', '1', '1', '0', '1'},
		[]byte{'0', '0', '0', '0', '1'},
	}
	fmt.Println(numIslands(grid))
}
