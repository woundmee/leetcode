package main

import "fmt"

func main() {
	res := guessNumber(10)
	fmt.Println(res)
}

// problem --> https://neetcode.io/problems/guess-number-higher-or-lower

// функция guess возврващет 0, -1 или 1, где 0 - совпадение, -1 - загаданная цифра меньше текущего, 1 - выше текущего.

func guess(num int) int // дается системой

func guessNumber(n int) int {

	left, right := 0, n

	for left <= right {

		mid := left + (right-left)/2
		res := guess(mid)

		switch {
		case res == 0:
			return mid
		case res == 1:
			left = mid + 1
		case res == -1:
			right = mid - 1
		}

	}

	return -1
}
