package main

import (
	"math"
)

// problem --> https://neetcode.io/problems/score-of-a-string/

// Input: s = "code"
// Output: 24

// c=99, o=111, d=100, e=101
// |111 - 99| + |100 - 111| + |101 - 100| = 24

func main() {

	scoreOfString("code")
	scoreOfString("neetcode")
}

func scoreOfString(s string) int {
	if len(s) < 2 {
		return 0
	}

	var res float64
	right := 1

	for left := range s {
		sub := math.Abs(float64(s[right]) - float64(s[left]))
		right++
		res += sub
		if right == len(s) {
			break
		}
	}
	return int(res)
}
