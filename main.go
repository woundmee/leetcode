package main

import "fmt"

// problem --> https://neetcode.io/problems/validate-parentheses/

func main() {
	fmt.Println(isValid("]"))
	fmt.Println(isValid("(){}}{"))
	fmt.Println(isValid("([{}])"))
	fmt.Println(isValid("[(])"))
	fmt.Println(isValid("[([]]]"))
	fmt.Println(isValid("()()"))
	fmt.Println(isValid("[[]"))

}

func isValid(s string) bool {
	// работа со стэком:
	// если скобка открывающая - кладем в стэк
	// далее смотрим, если приходит закрывающая скобка - сравниваем со значением в стэке (с конца слайса)
	// если не совпадают - return false
	// если совпадают - удаляем из стэка открывающую скобку

	// ( - 40
	// ) - 41
	// [ - 91
	// ] - 93
	// { - 123
	// } - 125

	if len(s) < 1 {
		return false
	}

	brackets := make([]byte, 0)

	if s[0] == 41 || s[0] == 93 || s[0] == 125 {
		return false
	}

	for i := range s {
		if s[i] == 40 || s[i] == 91 || s[i] == 123 {
			brackets = append(brackets, s[i])
		} else {
			if len(brackets) == 0 {
				return false
			}
			lastBracket := brackets[len(brackets)-1]

			if s[i]-1 != lastBracket && s[i]-2 != lastBracket {
				return false
			}
			brackets = brackets[:len(brackets)-1]

		}
	}

	if len(brackets) != 0 {
		return false
	}

	return true

}
