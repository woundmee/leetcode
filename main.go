package main

import (
	"strings"
	"unicode"
)

// problem:	https://neetcode.io/problems/decode-string/
// level:	medium

// task:	дана закондированная строка по шаблону k[string], где string необходимо повторить k-раз
//			пример: дана строка 2[a3[b]]c, на выходе: abbbabbbc

func main() {
	decodeString("2[a3[b]]c") // --> abbbabbbc
}

func decodeString(s string) string {

	letters := []string{}
	nums := []int{}

	currLetter := ""
	currNum := 0

	for _, v := range s {

		if unicode.IsLetter(v) {
			currLetter += string(v)
		} else if unicode.IsDigit(v) {

			// т.к. цифры могут идти рядом друг с другом, пример ..a12[b]..,
			// 12 он будет считать отдельно по цифрам, как 1, 2 и тд, а нам нужно получить 12,
			// чтобы повторить 12 раз, поэтому это простой способ до 99 сконвертить в десятичный формат
			currNum = currNum*10 + int(v-'0')
		} else {
			// если скобки
			switch v {
			case '[':
				letters = append(letters, currLetter)
				nums = append(nums, currNum)
				currLetter = ""
				currNum = 0

			case ']':
				k := pop(&nums)
				prev := pop(&letters)
				curr := currLetter
				currLetter = prev + strings.Repeat(curr, k)
			}
		}

	}

	return currLetter
}

func pop[T any](stack *[]T) T {
	last := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return last
}
