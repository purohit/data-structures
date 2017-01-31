package main

import "fmt"

func validUtf8(data []int) bool {
	for len(data) > 0 {
		// Found a 1-byte data piece
		if data[0]>>7 == 0 {
			data = data[1:]
			continue
		}
		// Found a 2- to 4-byte data piece
		var legal bool
		for k := 4; k >= 2; k-- {
			var mask int
			for j := uint(7); j >= 8-uint(k); j-- {
				mask |= 1 << j
			}
			if data[0]&mask == mask {
				if len(data) < k {
					return false
				}
				for i := 1; i <= k-1; i++ {
					if data[i]&(1<<7) != 1<<7 {
						return false
					}
				}
				data = data[k:]
				legal = true
				break
			}
		}
		if !legal {
			return false
		}
	}
	return true
}

func main() {
	data := []int{197, 130, 1}
	fmt.Println(validUtf8(data))
	data = []int{235, 140, 4}
	fmt.Println(validUtf8(data))
	// data = [197, 130, 1], which represents the octet sequence: 11000101 10000010 00000001.
	// Return true.
	// It is a valid utf-8 encoding for a 2-bytes character followed by a 1-byte character.
}
