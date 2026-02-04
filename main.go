package main

import (
	"fmt"
	"slices"
)

// problem -> https://neetcode.io/problems/top-k-elements-in-list/

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	for range k {
		max := slices.Max(nums)
		res = append(res, max)
		_ = slices.DeleteFunc(
			nums,
			func(v int) bool { return v == max },
		)
	}

	fmt.Println(res)
	return res
}

func main() {
	topKFrequent([]int{1, 2, 2, 3, 3, 3, 5, 7, 7}, 2)
	topKFrequent([]int{7, 7}, 1)
}
