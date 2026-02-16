package main

import (
	"strconv"
)

// problem --> https://neetcode.io/problems/baseball-game/

func main() {
	calPoints([]string{"1", "2", "+", "C", "5", "D"})
}

// in: [1 2 + C 5 D]
// out: 18

// explantation:
// 1 - add
// 2 - add   --> [1 2]
// + - суммирование 2х предыдущих: 1+2  -> [1 2 3]
// C - удалить предыдущий результат  -> [1 2]
// 5 - добавление  -> [1 2 5]
// D - умножение 2х предыдущих: 2*5  -> [1 2 5 10]

// out: сумма всех: 1+2+5+10 = 18

// ================= //
// 1 2

func calPoints(operations []string) int {

	stack := make([]int, 0, len(operations))

	for i := range operations {
		if digit, err := strconv.Atoi(operations[i]); err == nil {
			stack = append(stack, digit)
			continue
		}

		switch operations[i] {
		case "+":
			// note: нужно условие проверки на len(stack) >= 2 ?!
			n := len(stack)
			addition := stack[n-1] + stack[n-2]
			stack = append(stack, addition)
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		}
	}

	var sum int
	for i := range stack {
		sum += stack[i]
	}

	return sum
}
