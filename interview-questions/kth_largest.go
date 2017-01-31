package main

import (
	"container/heap"
	"fmt"
)

type inth []int

func (h inth) Less(i, j int) bool { return h[i] < h[j] }
func (h inth) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h inth) Len() int           { return len(h) }
func (h *inth) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *inth) Pop() interface{} {
	last := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return last
}

func findKthLargest(nums []int, k int) int {
	var h inth
	heap.Init(&h)
	for i, n := range nums {
		if i < k {
			heap.Push(&h, n)
		} else if n >= h[0] {
			heap.Pop(&h)
			heap.Push(&h, n)
		}
	}
	return h[0]
}

func main() {
	type t struct {
		nums []int
		k    int
	}
	for _, i := range []t{
		t{[]int{1, 2, 3}, 1},
		t{[]int{1, 2, 3}, 2},
		t{[]int{1, 2, 3}, 3},
		t{[]int{5, 2, 1, -16, 15, 3, 4, 80, 11, 11}, 1},
		t{[]int{5, 2, 1, -16, 15, 3, 4, 80, 11, 11}, 3},
		t{[]int{5, 2, 1, -16, 15, 3, 4, 80, 11, 11}, 4},
		t{[]int{5, 2, 1, -16, 15, 3, 4, 80, 11, 11}, 5},
	} {
		fmt.Println(findKthLargest(i.nums, i.k))
	}
}
