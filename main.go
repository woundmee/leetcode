package main

// problem --> https://neetcode.io/problems/move-zeroes/

func main() {
	moveZeroes([]int{0, 0, 1, 2, 0, 5})
	moveZeroes([]int{1, 0, 2, 0, 0, 3, 4})

}

// in:	0,0,1,2,0,5
// out:	1,2,5,0,0,0

// note: задача двигать нули в конец

func moveZeroes(nums []int) {
	var left int
	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
