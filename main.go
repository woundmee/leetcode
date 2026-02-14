package main

import (
	"fmt"
	"strings"
	"unicode"
)

// problem --> https://neetcode.io/problems/is-palindrome/

func main() {

	// res := isPalindrome("Was it a car or a cat I saw?")
	res := isPalindrome("tab a cat")
	fmt.Println(res)

}

// 2 указателя

// Input: s = "Was it a car or a cat I saw?"
// Output: true

func isPalindrome(s string) bool {

	left, right := 0, len(s)-1
	r := []rune(s)

	for right > left {

		// letel?
		// ^   ^

		if !unicode.IsLetter(r[left]) && !unicode.IsDigit(r[left]) {
			left++
			continue
		} else if !unicode.IsLetter(r[right]) && !unicode.IsDigit(r[right]) {
			right--
			continue
		}

		fmt.Println(
			string(r[left]),
			string(r[right]),
		)

		if strings.EqualFold(string(r[left]), string(r[right])) {
			left++
			right--
		} else {
			return false
		}
	}

	return true

}
