package main

import (
	"fmt"
	"slices"
)

// problem -> https://neetcode.io/problems/anagram-groups/
// задача: сгруппировать анаграммы

// Input: strs = ["act","pots","tops","cat","stop","hat"]
// Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

func main() {
	// groupAnagrams([]string{"eat", "cat", "tea", "koko", "okko"})
	groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"})
}

func groupAnagrams(strs []string) [][]string {

	var res [][]string

	if len(strs) < 1 {
		return res
	}

	m := make(map[string][]string)

	for _, v := range strs {
		s := sortString(v)
		if _, ok := m[s]; ok {
			m[s] = append(m[s], v)
		} else {
			m[s] = append(m[s], v)
		}
	}

	for _, v := range m {
		res = append(res, v)
	}

	fmt.Println(res)
	return res
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}
