package main

// problem -> https://neetcode.io/problems/valid-palindrome-ii/
// задача: проверить текст на палиндром.
// У тебя есть 1 попытка убрать букву (с левой или правой стороны)

// example: "addba" (можно убрать предпоследнюю букву и слово == палиндром)

func main() {
	// validPalindrome("acdccba")  // false
	validPalindrome("abbda") // true
	// validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga")  // true
}

func validPalindrome(s string) bool {

	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {

			l := ispal(s, left+1, right)
			r := ispal(s, left, right-1)
			return l || r

		} else {
			left++
			right--
		}
	}
	return true
}

func ispal(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
