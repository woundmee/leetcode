package main

import "fmt"

// problem --> https://neetcode.io/problems/merge-strings-alternately/

func main() {

	fmt.Println(
		mergeAlternately("abc", "xyz"),
		mergeAlternately("ab", "abbxxc"),
	)

}

// Input: word1 = "abc", word2 = "xyz"
// Output: "axbycz"

func mergeAlternately(word1 string, word2 string) string {

	var minLen int
	var isWord1Min bool
	if len(word1) > len(word2) {
		minLen = len(word2)
		isWord1Min = false
	} else {
		minLen = len(word1)
		isWord1Min = true
	}

	var res string

	for i := 0; i <= minLen-1; i++ {
		res += string(word1[i]) + string(word2[i])
		if i == minLen-1 {
			if isWord1Min {
				res += word2[i+1:]
			} else {
				res += word1[i+1:]
			}
		}
	}

	return res
}
