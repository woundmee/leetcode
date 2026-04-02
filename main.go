package main

import "fmt"

// problem:	https://neetcode.io/problems/merge-sorted-array/
// level:	easy

// task:	расставить цифры из nums2 в nums1 по порядку (на свои места). Массивы отсортированы.

func main() {
	merge([]int{10, 20, 20, 40, 0, 0}, 4, []int{1, 2}, 2)
	// merge([]int{0, 0}, 0, []int{1, 2}, 2)
	// merge([]int{1, 2, 3, 4, 0, 0}, 4, []int{5, 6}, 2)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j, l := m-1, n-1, m+n-1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[l] = nums1[i]
			i--
		} else {
			nums1[l] = nums2[j]
			j--
		}
		l--
	}
	fmt.Println(nums1)
}
