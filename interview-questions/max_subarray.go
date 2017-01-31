package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best := nums[0]
	prevBest := nums[0]
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur > cur+prevBest {
			prevBest = cur
		} else {
			prevBest += cur
		}
		if prevBest > best {
			best = prevBest
		}
	}
	return best
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1, -1, -1, 5, -6, 20}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-11}))
}
