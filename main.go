package main

import "fmt"

// problem --> https://neetcode.io/problems/replace-elements-with-greatest-element-on-right-side/

func main() {
	replaceElements([]int{2, 4, 5, 3, 1, 2})
}

// Input: arr = [2,4,5,3,1,2]
// Output: [5,5,3,2,2,-1]

// заменить текущеий элемент самым большим элементом справа, а последний заменить на -1.
func replaceElements(arr []int) []int {

	maxRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		curr := arr[i]
		arr[i] = maxRight

		if curr > maxRight {
			maxRight = curr
		}
	}

	fmt.Println(arr)
	return arr
}
