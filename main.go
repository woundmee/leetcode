package main

import "fmt"

// problem -> https://leetcode.com/problems/two-sum/description/

func main() {
	s := []int{3, 2, 3}
	fmt.Println(twoSum(s, 6))
}

// nums = [2, 7, 1, 5], target = 9, out: [0, 1]
func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := range i {
			if nums[j]+nums[i] == target {
				return []int{j, i}
			}
		}
	}
	return []int{}
}
