package main

import (
	"fmt"
)

// problem:	https://neetcode.io/problems/first-missing-positive/
// level:	hard

// task:	вернуть первый отсутствующий элемент в массиве (старт с 1)

func main() {
	// firstMissingPositive([]int{-2, -1, 0})
	// firstMissingPositive([]int{1, 2, 4})
	// firstMissingPositive([]int{1, 2, 4, 5, 6, 3, 1})
	// firstMissingPositive([]int{1})
	// firstMissingPositive([]int{1, 2, 3, 4, 5})
	firstMissingPositive([]int{5, 4, 3, 2, 1})
}

func firstMissingPositive(nums []int) int {

	m := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			fmt.Println(i)
			return i
		}
	}

	return len(nums) + 1
}
