package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type nodeHeap []*ListNode

func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h nodeHeap) Len() int           { return len(h) }
func (h *nodeHeap) Push(x interface{}) {
	(*h) = append(*h, x.(*ListNode))
}
func (h *nodeHeap) Pop() interface{} {
	max := len(*h) - 1
	last := (*h)[max]
	(*h) = (*h)[:max]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head, cur *ListNode
	var h nodeHeap
	for _, l := range lists {
		if l == nil {
			continue
		}
		heap.Push(&h, l)
	}
	// Consume until the heap is empty
	for h.Len() > 0 {
		l := heap.Pop(&h).(*ListNode)
		node := &ListNode{Val: l.Val}
		if head == nil {
			head = node
			cur = node
		} else {
			cur.Next = node
			cur = cur.Next
		}
		if l.Next != nil {
			heap.Push(&h, l.Next)
		}
	}
	return head
}

func main() {
	a := &ListNode{Val: 1}
	a.Next = &ListNode{Val: 2}
	a.Next.Next = &ListNode{Val: 5}
	a.Next.Next.Next = &ListNode{Val: 6}

	b := &ListNode{Val: -1}
	b.Next = &ListNode{Val: 0}
	b.Next.Next = &ListNode{Val: 3}
	b.Next.Next.Next = &ListNode{Val: 17}
	b.Next.Next.Next.Next = &ListNode{Val: 18}

	var c *ListNode

	r := mergeKLists([]*ListNode{a, b, c})
	for r != nil {
		fmt.Printf("%d ", r.Val)
		r = r.Next
	}
}
