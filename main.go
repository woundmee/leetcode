package main

import "fmt"

// problem --> https://leetcode.com/problems/minimum-string-length-after-removing-substrings/

func main() {
	minLength("ABFCACDB")
	minLength("ACBBD")
	minLength("BJKDKABJ")
}

func minLength(s string) int {

	stack := []rune{}
	// stack := make([]rune, len(s)/2)
	for _, curr := range s {
		if len(stack) > 0 {
			last := stack[len(stack)-1]
			if (last == 'A' && curr == 'B') || (last == 'C' && curr == 'D') {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, curr)
	}

	fmt.Println(string(stack), "-", len(stack))
	return len(stack)
}
