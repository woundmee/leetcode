package main

import (
	"fmt"
)

// problem -> https://leetcode.com/problems/single-number/

func main() {
	// s := []int{4, 1, 2, 1, 2}
	fmt.Println(isAnagram(
		"rellepg",
		"speller",
	),
	)

}

// Входные данные: nums = [4,1,2,1,2]
// Вывод: 4

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
