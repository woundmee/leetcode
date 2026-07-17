package main

// problem:	https://neetcode.io/problems/remove-duplicates-from-sorted-array/
// level:	easy

// task:	...

func main() {
	removeDuplicates([]int{2, 10, 10, 30, 30, 30})
	removeDuplicates([]int{1, 2, 2, 3, 4, 5, 5, 6, 7, 7, 8})
}

func removeDuplicates(nums []int) int {
	l, r := 0, 1

	for r != len(nums) {
		if nums[l] == nums[r] {
			r++
		} else if nums[r] > nums[l] {
			nums[l+1] = nums[r]
			l++
			r++
		}
	}
	return len(nums[:l+1])
}
