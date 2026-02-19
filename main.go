package main

import (
	"unicode"
)

// problem --> https://leetcode.com/problems/clear-digits/

func main() {
	clearDigits("abc34")
}

// abc34
// удаляем цифру удаляем ближайщий к нему символ
// out: a

// abc34
//     ^
// stack: [a]

func clearDigits(s string) string {

	stack := []rune{}
	for _, v := range s {
		if unicode.IsLetter(v) {
			stack = append(stack, v)
			continue
		}

		if len(stack) != 0 && unicode.IsDigit(v) {
			stack = stack[:len(stack)-1]
		}
	}

	var res string
	for _, v := range stack {
		res += string(v)
	}

	return res

}
