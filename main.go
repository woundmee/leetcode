package main

import (
	"fmt"
	"strconv"
)

// problem:	https://neetcode.io/problems/evaluate-reverse-polish-notation
// level:	medium

// task:	Вам дан массив строковых токенов, представляющих допустимое
// 			арифметическое выражение в обратной польской нотации.

func main() {
	res := evalRPN([]string{"1", "2", "+", "3", "*", "4", "-"})
	fmt.Println(res)
}

// ["1","2","+","3","*","4","-"]
// out = ((1+2)*3)-4=9-4 = 5

// tokens=["4","13","5","/","+"]
// out = 6

func evalRPN(tokens []string) int {

	stack := make([]int, 0)

	for _, v := range tokens {
		dig, err := strconv.Atoi(v)

		if err == nil {
			stack = append(stack, dig)
		} else {
			operations(&stack, v)
		}

	}

	return stack[0]
}

func operations(stack *[]int, oper string) {
	switch oper {
	case "+":
		last := pop(stack)
		penult := pop(stack)
		*stack = append(*stack, last+penult)
	case "-":
		last := pop(stack)
		penult := pop(stack)
		*stack = append(*stack, penult-last)
	case "*":
		last := pop(stack)
		penult := pop(stack)
		*stack = append(*stack, last*penult)
	case "/":
		last := pop(stack)
		penult := pop(stack)
		*stack = append(*stack, penult/last)
	}
	fmt.Println(*stack)
}

// через указатель
func pop(stack *[]int) int {
	s := *stack
	removed := s[len(*stack)-1]
	*stack = s[:len(*stack)-1]
	return removed
}

// // копированием
// func pop(stack []int) (int, []int) {
// 	last := stack[len(stack)-1]
// 	newStack := stack[:len(stack)-1]
// 	return last, newStack
// }
