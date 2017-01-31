package main

import (
	"bytes"
	"fmt"
)

func decodeString(s string) string {
	var counts []int
	var results []string
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			// Consume digits
			var n int
			for ; s[i] >= '0' && s[i] <= '9'; i++ {
				n = n*10 + int(s[i]-'0')
			}
			// [ is skipped in next iteration
			counts = append(counts, n)
			results = append(results, result)
			result = ""
			continue
		} else if s[i] == ']' {
			tmp := bytes.NewBufferString(results[len(results)-1])
			results = results[:len(results)-1]
			for c := counts[len(counts)-1]; c > 0; c-- {
				tmp.WriteString(result)
			}
			counts = counts[:len(counts)-1]
			result = tmp.String()
			continue
		}
		result += string([]byte{s[i]})
	}
	return result
}

func main() {
	fmt.Println(decodeString("02[abc]3[x4[y]]zvv"))
	fmt.Println(decodeString(""))
	fmt.Println(decodeString("101[lee]"))
}
