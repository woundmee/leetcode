package main

import "fmt"

func main() {
	res := searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Println(res)
}

// problem --> https://neetcode.io/problems/search-insert-position/

// дается массив и target. Нужно вернуть индекс таргета, а если таркет отсутствует в списке,
// то нужно вернут индекс, куда этот target был бы вставлен.

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		}
	}
	return left
}
