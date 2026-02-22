package main

// problem --> https://leetcode.com/problems/apply-operations-to-an-array/

func main() {
	// applyOperations([]int{1, 2, 2, 1, 1, 0})
	applyOperations([]int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272})
}

// in:	[1,2,2,1,1,0]
// out:	[1,4,2,0,0,0]

// если соседние элементы равны (i == i+1), тогда i умножаем на 2 и записываем
// вместо него новый результат, а в i+1 записываем 0.
// в конце просто сдвигаем все нули в конец

func applyOperations(nums []int) []int {

	// step #1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	// step #2
	left := 0
	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	return nums
}
