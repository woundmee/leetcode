package main

import "fmt"

// problem --> https://neetcode.io/problems/is-subsequence/

func main() {
	// isSubsequence("node", "neetcode")
	isSubsequence("abc", "aabbcc")
}

// Input: s = "node", t = "neetcode"
// Output: true

// если строка 's' является последовательностью строки 't', то вернуть true
// строка последовательна, если все буквы из 's' идут попорядку в 't', пример:
// s=asec, t=asneg, в 't' есть только ase, которые идут попорядку

func isSubsequence(s string, t string) bool {

	// s=abc   t=aabbcc
	//    ^       ^
	// res=a

	if len(s) == 0 {
		return true
	}

	leftOne := 0
	res := make([]byte, len(s))
	for i := range t {
		if s[leftOne] == t[i] && res[leftOne] != t[i] {
			res[leftOne] = t[i]
			leftOne++
			if string(res) == s {
				return true
			}
		}
	}

	fmt.Println(string(res))
	// fmt.Println(s == res)
	return s == string(res)
}
