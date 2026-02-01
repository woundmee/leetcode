package main

// problem -> https://leetcode.com/problems/move-zeroes/

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

}

// Входные данные: nums = [0,1,0,3,12]
// Выходные данные: [1,3,12,0,0]

func moveZeroes(nums []int) {

	if len(nums) <= 1 {
		return
	}

	last := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[last], nums[i] = nums[i], nums[last]
			last++ // сдвигаю до следующего НЕнулевого элемента
		}
	}
}
