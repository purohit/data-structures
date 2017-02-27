package main

import "fmt"

func findMaxLength(nums []int) int {
	counts := map[int]int{}
	count, max := 0, 0
	for i, n := range nums {
		if n == 1 {
			count++
		} else {
			count--
		}
		if i > 0 && count == 0 && i+1 > max {
			max = i + 1
		}
		if _, ok := counts[count]; !ok {
			counts[count] = i
		} else if dist := i - counts[count]; dist > max {
			max = dist
		}
	}
	return max
}

func main() {
	tests := [][]int{
		[]int{},
		[]int{0, 1},
		[]int{1, 0, 1, 0, 1},
		[]int{1, 0, 0, 0, 1, 1, 1, 1},
		[]int{0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1},
		[]int{0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1},
		[]int{0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1},
	}
	for _, t := range tests {
		fmt.Println(findMaxLength(t))
	}
}
