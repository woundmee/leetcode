package main

import "fmt"

// problem --> https://neetcode.io/problems/reverse-string/

func main() {

	reverseString([]byte("neet"))

}

// ИСПОЛЬЗОВАЛ алгоритм "2 указателя"

// перевернуть исходный массив - O(1)
// Input: s = ["n","e","e","t"]
// Output: ["t","e","e","n"]

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	fmt.Println(string(s))

	for right > left {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	fmt.Println(string(s))
}
