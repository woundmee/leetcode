package main

import "fmt"

// problem -> https://leetcode.com/problems/two-sum/description/

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println(twoSum(s, 5))
}

// nums = [2, 7, 1, 5], target = 9, out: [0, 1]
func twoSum(nums []int, target int) []int {
	for i := 1; i <= len(nums); i++ {
		if nums[i]+nums[i-1] == target {
			return []int{i - 1, i}
		}
	}
	return []int{}
}
