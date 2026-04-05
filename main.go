package main

import "fmt"

// problem:
// level:	medium

// task:	дается отсортированный массив и K-тый элемент.
// 			Вернуть массив индексов, числа которых в сумме дают К. Индексы вернуться, начиная с 1, а не 0.

//			[2 6 7 8 10 40 50], target=15
//			out: [3 4] (3-ий индекс == 7, 4-ый == 8), 7+8=15

func main() {
	twoSum([]int{3, 6, 10, 12, 15, 17, 20}, 25)
	// twoSum([]int{1, 2, 3, 4}, 3)
}

func twoSum(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		switch {
		case sum == target:
			fmt.Println([]int{left + 1, right + 1})
			return []int{left + 1, right + 1}
		case sum > target:
			right--
		case sum < target:
			left++
		}
	}

	return []int{}
}
