package main

import "fmt"

func finalLen(S string, K int) int {
	good := 0
	for i := range S {
		if S[i] != '-' {
			good++
		}
	}
	groups := good / K
	if good%K > 0 {
		groups++
	}
	return good + groups - 1
}

func licenseKeyFormatting(S string, K int) string {
	formatted := make([]byte, finalLen(S, K))
	for i, fi := len(S)-1, len(formatted)-1; fi >= 0; {
		if S[i] == '-' {
			i--
			continue
		}
		fromEnd := len(formatted) - fi
		if fromEnd%(K+1) == 0 {
			formatted[fi] = '-'
			fi--
			continue
		}
		formatted[fi] = S[i]
		if formatted[fi] >= 'a' && formatted[fi] <= 'z' {
			formatted[fi] -= ('a' - 'A') // Uppercase it
		}
		i, fi = i-1, fi-1
	}
	return string(formatted)
}
func main() {
	x := licenseKeyFormatting("2-4A0r7-4k", 3)
	fmt.Println(x)
}
