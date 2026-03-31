package main

// problem -> https://neetcode.io/problems/sort-an-array/
// задача: отсортировать заданный массив. Использовал метод QuickSort

func main() {
	sortArray([]int{6, 2, 4, 1, 3, 5})
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
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
