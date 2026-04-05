package main

import "fmt"

// problem:	https://neetcode.io/problems/rotate-array/
// level:	medium

// task: 	дан массив, необходимо его перевернуть k-раз. Если к > len(nums), то k %= len(nums)
// 			Input: nums = [1,2,3,4,5,6,7,8], k = 4
// 			Output: [5,6,7,8,1,2,3,4]

// 			Пояснение:
// 			поворот на 1 шаг вправо: [8,1,2,3,4,5,6,7]
// 			поворот на 2 шага вправо: [7,8,1,2,3,4,5,6]
// 			поворот на 3 шага вправо: [6,7,8,1,2,3,4,5]
// 			поворот на 4 шага вправо: [5,6,7,8,1,2,3,4]

func main() {
	// rotate([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4)
	rotate([]int{1, 2, 3, 4, 5}, 7)
}

// это решение является более оптимальным,
// чем просто написать: nums = append(nums[l:], nums[:l]...)
// решение в одну строку, правда создается новый массив, а решение ниже оптимально тем,
// что меняю исходный на месте (in-place), не создавая новый массив (как при append)

func rotate(nums []int, k int) {
	if len(nums) <= k {
		k %= len(nums)
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	fmt.Println(nums)
}

func reverse(nums []int, l, r int) {
	for l <= r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
