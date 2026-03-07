package main

import "fmt"

func main() {
	res := search([]int{-1, 0, 2, 4, 6, 8}, 4)
	fmt.Println(res)
}

// problem --> https://neetcode.io/problems/binary-search/

// необходимо реализовать простую функцию для бинарного поиска
// если target отсутствует, вернуть -1

func search(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // note: запомни эту формулу

		switch {
		case nums[mid] == target:
			return nums[mid]
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		}
	}

	return -1
}
