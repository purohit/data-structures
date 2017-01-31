package main

import "math/rand"
import "fmt"

type RandomizedSet struct {
	vals map[int]int
	pos  []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		vals: make(map[int]int, 0),
		pos:  make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.vals[val]; ok {
		return false
	}
	this.pos = append(this.pos, val)
	this.vals[val] = len(this.pos) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.vals[val]; !ok {
		return false
	}
	// If it's not at the end index, make it so
	if curIdx, lastIdx := this.vals[val], len(this.pos)-1; curIdx != lastIdx {
		this.pos[lastIdx], this.pos[curIdx] = this.pos[curIdx], this.pos[lastIdx]
		this.vals[this.pos[curIdx]] = curIdx
	}
	delete(this.vals, val)
	this.pos = this.pos[:len(this.pos)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.pos[rand.Intn(len(this.pos))]
}

func main() {
	r := Constructor()
	fmt.Println(r.Insert(1))
	fmt.Println(r.Remove(2))
	fmt.Println(r.Insert(2))
	fmt.Println(r.GetRandom())
	fmt.Println(r.Remove(1))
	fmt.Println(r.Insert(2))
	fmt.Println(r.GetRandom())
}
