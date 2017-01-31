package main

import "fmt"

type rotator string

func (r rotator) key() string {
	key := ""
	for i := 1; i < len(r); i++ {
		dist := (r[i] - r[i-1] + 26) % 26
		key += fmt.Sprintf("|%d", dist)
	}
	return key
}

func group(toGroup []rotator) [][]rotator {
	rots := make(map[string][]rotator)
	for _, g := range toGroup {
		k := g.key()
		if _, ok := rots[k]; !ok {
			rots[k] = make([]rotator, 0)
		}
		rots[k] = append(rots[k], g)
	}
	result := make([][]rotator, 0)
	for _, v := range rots {
		result = append(result, v)
	}
	return result
}

func main() {
	tests := []rotator{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z", "baba", "azaz", "fefe"}
	fmt.Println(group(tests))
}
