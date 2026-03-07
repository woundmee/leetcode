package main

// problem --> https://leetcode.com/problems/backspace-string-compare/

func main() {
	// backspaceCompare("abcde##", "abcde##")
	backspaceCompare("ab##", "c#d#")
}

// in -->  s1: "ab#c"  s2: "ah#c"
// out -->  true

// цель: удалить символы '#' и те, которые идут до них

func backspaceCompare(s string, t string) bool {
	return checkSymbol(s) == checkSymbol(t)
}

func checkSymbol(str string) string {
	stack := make([]byte, 0, len(str))

	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			stack = append(stack, str[i])
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
