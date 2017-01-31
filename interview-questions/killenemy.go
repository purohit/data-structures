package main

import (
	"fmt"
)

type killed struct{ h, v int }

func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	killedByPoint := make([][]killed, len(grid))
	for i := range grid {
		killedByPoint[i] = make([]killed, len(grid[i]))
	}
	max := 0
	// Calculate horizontal kill radius, then run through again
	for r := range grid {
		k := 0
		for left, right := 0, 0; right < len(grid[r]); right++ {
			if grid[r][right] == 'E' {
				k++
			}
			if grid[r][right] == 'W' || right == len(grid[r])-1 {
				for ; left <= right; left++ {
					killedByPoint[r][left].h = k
				}
				k = 0
			}
		}
	}
	// Vertical kill radius; same thing here in the other direction.
	cols := len(grid[0])
	for c := 0; c < cols; c++ {
		k := 0
		for top, bottom := 0, 0; bottom < len(grid); bottom++ {
			if grid[bottom][c] == 'E' {
				k++
			}
			if grid[bottom][c] == 'W' || bottom == len(grid)-1 {
				for ; top <= bottom; top++ {
					here := killedByPoint[top][c]
					here.v = k
					if grid[top][c] == '0' {
						if here.h+here.v > max {
							max = here.h + here.v
						}
					}
				}
				k = 0
			}
		}
	}
	return max
}

func main() {
	grid := [][]byte{
		[]byte{'0', 'E', '0', '0'},
		[]byte{'E', '0', 'W', 'E'},
		[]byte{'0', 'E', '0', '0'},
	}
	fmt.Println(maxKilledEnemies(grid))
	fmt.Println(maxKilledEnemies([][]byte{}))
}
