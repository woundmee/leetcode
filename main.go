package main

import "fmt"

// problem -> https://leetcode.com/problems/roman-to-integer/

func main() {
	s := []int{1, 1, 2}
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i <= len(nums)-1; i++ {
		// если текущий элемент != предыдущему - мы нашли новый элемент!
		if nums[i] != nums[i-1] {
			nums[k] = nums[i] // передвигаем его в начало
			k++
		}
	}

	return k
}
