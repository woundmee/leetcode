package main

import (
	"slices"
)

// problem -> https://neetcode.io/problems/longest-consecutive-sequence/
// задача: дан массив, вернуть самую длинную последовательность, которому можно составить в массиве
// пример:
// in: [0,3,2,5,4,6,1,1]
// out: 7
// максимальная последовательность: 0 1 2 3 4 5 6 (7 раз)

func main() {
	longestConsecutive([]int{2, 20, 4, 10, 3, 4, 5})
	// longestConsecutive([]int{0, 3, 2, 5, 4, 6, 1, 1})
	// longestConsecutive([]int{2, 20, 4, 10, 3, 4, 5})
	longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})
}

func longestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	slices.Sort(nums)

	// 9,1,4,7,3,-1,0,5,8,-1,6
	// -1 -1 0 1 3 4 5 6 7 8 9
	//                       ^
	// count=9

	var count int = 1
	counts := []int{}

	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i]-1 != nums[i-1] {
			counts = append(counts, count)
			count = 0
		}
		count++
	}

	counts = append(counts, count)

	println(slices.Max(counts))
	return slices.Max(counts)
}
