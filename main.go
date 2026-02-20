package main

// problem --> https://leetcode.com/problems/make-the-string-great/

func main() {
	makeGood("leEeetcode")
	makeGood("abBAcC")
}

// удаляем парные буквы разных регистров, которые находятся вместе
// пример: aA, Aa
// leEeetcode  --> eE - удаляем --> leetcode

func makeGood(s string) string {
	stack := make([]byte, 0)

	for i := range s {
		if len(stack) != 0 {
			last := stack[len(stack)-1]
			// [a b] B
			if last-32 == s[i] || last+32 == s[i] {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
