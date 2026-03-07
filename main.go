package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"bat", "bag", "bank", "band"}))
}

// problem --> https://neetcode.io/problems/longest-common-prefix/

// Input: strs = ["bat","bag","bank","band"]
// Output: "ba"

// найти самое длинное вхождение букв. Если хоть в одном слове будет несовпадение, вхождение считается по минимому:
// пример: ["tat", "tato", "task", "tatpool"]  --> макс вхождение = ta

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// ["bat","bag","bank","band"]
	//   ^

	for i := range strs[0] {
		char := strs[0][i]

		for j := range len(strs) {
			if i == len(strs[j]) || char != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
