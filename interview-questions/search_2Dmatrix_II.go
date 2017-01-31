package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for r, c := 0, len(matrix[0])-1; r < len(matrix) && c >= 0; {
		if matrix[r][c] == target {
			return true
		} else if target < matrix[r][c] {
			c--
		} else if target > matrix[r][c] {
			r++
		}
	}
	return false
}

func main() {
	m := [][]int{
		[]int{1, 4, 7, 11, 15, 16},
		[]int{2, 5, 8, 12, 19, 19},
		[]int{3, 6, 9, 16, 22, 22},
		[]int{10, 13, 14, 17, 24, 44},
		[]int{18, 21, 23, 26, 30, 50},
	}
	fmt.Println(searchMatrix(m, 5))
	fmt.Println(searchMatrix(m, 21))
	fmt.Println(searchMatrix(m, 33))
	fmt.Println(searchMatrix(m, 20))
	fmt.Println(searchMatrix(m, -5))
	fmt.Println(searchMatrix(m, 50))
}
