package main

import (
	"fmt"
)

func main() {
	// res := findPeakElement([]int{1, 2, 3, 1})
	// res := findPeakElement([]int{1, 2, 1, 3, 4, 5, 0})
	// res := findPeakElement([]int{1, 2, 3})
	// res := findPeakElement([]int{1, 2})
	res := findPeakElement([]int{3, 2, 1})
	fmt.Println(res)
}

// problem --> https://neetcode.io/problems/find-peak-element/

// in: [1,2,1,3,4,5,0]
// out: 5 (index)

// найти пиковый элемент в массиве и вернуть его индекс.
// пиковый - это когда элемент слева и справа от него меньше него самого

func findPeakElement(nums []int) int {

	left, right := 0, len(nums)-1

	// [1 2 3 1]
	//    ^

	for left < right {
		mid := left + (right-left)/2
		if mid > 0 && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
