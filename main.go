package main

import "fmt"

// problem:	https://neetcode.io/problems/sort-colors/
// task:	отсортировать "цвета" (в виде цифр). Применил тот же алгоритм - QuickSort.

func main() {
	sortColors([]int{1, 0, 1, 2})
}

func sortColors(nums []int) {
	fmt.Println(nums)
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := nums[(left+right)/2]
	i, j := left, right

	for i <= j {
		for pivot > nums[i] {
			i++
		}
		for pivot < nums[j] {
			j--
		}

		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	quickSort(nums, left, j)
	quickSort(nums, i, right)
}
