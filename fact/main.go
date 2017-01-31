package main

import "fmt"

func factor(n int) ([][]int, error) {
	if n < 2 {
		return nil, fmt.Errorf("Can't compute factoring of small numbers.")
	}
	results := make([][]int, 0)
	factorA(n, n, &results, nil)
	return results, nil
}

func factorA(orig, cur int, results *[][]int, oldSet []int) {
	for i := 2; i <= orig/2; i++ {
		if cur%i != 0 {
			continue
		}
		var curSet []int
		if oldSet != nil {
			curSet = make([]int, len(oldSet))
			copy(curSet, oldSet)
		}
		curSet = append(curSet, i)
		if cur/i == 1 { // Found a hit
			*results = append(*results, curSet)
			return
		}
		factorA(orig, cur/i, results, curSet)
	}
}

func main() {
	// Can use to find primes?
	for i := 2; i < 400; i++ {
		results, err := factor(i)
		if err != nil {
			fmt.Println(err)
		}
		if len(results) == 0 {
			fmt.Println(i, results)
		}
	}
}
