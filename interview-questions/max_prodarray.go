package main

import "fmt"

func min(xs ...int) int {
	min := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < min {
			min = xs[i]
		}
	}
	return min
}

func max(xs ...int) int {
	max := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > max {
			max = xs[i]
		}
	}
	return max
}

func maxProdArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best, prevMax, prevMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		localMax := max(prevMax*cur, prevMin*cur, cur)
		localMin := min(prevMax*cur, prevMin*cur, cur)
		best = max(localMax, best)
		prevMax, prevMin = localMax, localMin
	}
	return best
}

func main() {
	fmt.Println(maxProdArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxProdArray([]int{2, 3, -2, 4}))
	fmt.Println(maxProdArray([]int{-2}))
	fmt.Println(maxProdArray([]int{-2, 3}))
	fmt.Println(maxProdArray([]int{1, 4, -5, 2, -6, 7}))
	fmt.Println(maxProdArray([]int{-1, 2, 5, 7, 0}))
}
