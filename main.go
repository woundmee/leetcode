package main

import (
	"strings"
)

// problem --> https://leetcode.com/problems/reverse-words-in-a-string-iii/

func main() {
	reverseWords("Let's take LeetCode contest")
}

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// перевернуть каждое слово так, чтобы сохранились символы, проблемы.
// решается 2-умя указателями.

// ok: сохранить в telegram

func reverseWords(s string) string {

	words := strings.Fields(s)
	res := []byte{}

	for i := range words {
		right := len(words[i]) - 1
		for right >= 0 {
			res = append(res, words[i][right])
			right--
		}
		res = append(res, ' ')
	}

	return string(res[:len(res)-1])
}
