package main

import (
	"fmt"
)

func updateMax(l, r int, s []rune, longest *int) {
	for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
		if s[l] != s[r] {
			break
		}
		if cur := r - l + 1; cur > *longest {
			*longest = cur
		}
	}
}

func longestPalindrome(s string) int {
	var longest int
	for i, r := 0, []rune(s); i < len(r); i++ {
		updateMax(i, i, r, &longest)
		if i > 0 {
			updateMax(i-1, i, r, &longest)
		}
	}
	return longest
}

func main() {
	for _, s := range []string{
		"",
		"abracadabra",
		"awagonwheel",
		"nutellalle",
		"a你好b好你lolwhat",
		"amanaplanacanalpanama",
	} {
		fmt.Println(len(s), longestPalindrome(s))
	}
}
