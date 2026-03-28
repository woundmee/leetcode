package main

import "fmt"

// problem -> https://neetcode.io/problems/products-of-array-discluding-self/
// задача: вернуть массив, где каждый n[i] равна сумме остальных элементов массива (не включая самого себя).

// Input: nums = [1,2,4,6]
// Output: [48,24,12,8]

// example: input: [1 2 3 4]
// output = [2*3*4, 1*3*4, 1*2*4, 1*2*3] = [24, 12, 8, 6]

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
	// productExceptSelf([]int{-1, 0, 1, 2, 3})
}

// O(n)
func productExceptSelf(nums []int) []int {

	if len(nums) < 1 {
		return []int{}
	}

	res := make([]int, len(nums))

	left := 1
	for i := range nums {
		res[i] = left
		left *= nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	fmt.Println(res)

	return res
}

// ===== O(n log n) ====== //
// ====================== //
// func productExceptSelf(nums []int) []int {
// 	var res []int
// 	numscopy := slices.Clone(nums)

// 	for i := range nums {
// 		var sum int = 1
// 		for j := range numscopy {
// 			if i == j {
// 				continue
// 			}
// 			sum *= numscopy[j]
// 		}
// 		res = append(res, sum)
// 	}

// 	fmt.Println(res)
// 	return res
// }
