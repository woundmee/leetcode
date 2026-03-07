package main

import (
	"fmt"
)

func main() {
	res := isPerfectSquare(15)
	fmt.Println(res)
}

// problem --> https://neetcode.io/problems/valid-perfect-square/

// дается положительно число, нужно вернуть полный ее квадрат.
// Полный квадрат - это произведение числа на саму себя, например, полный квадрат для 16 - это 4, потому 4*4=16
// для 15 полного квадрата нет, поэтому вернем false
// нельзя использовать встроенный math.sqrt()

func isPerfectSquare(num int) bool {

	left, right := 0, num

	for left <= right {
		mid := left + (right-left)/2

		square := mid * mid
		fmt.Println("square =", square)

		if square == num {
			return true
		} else if square < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
