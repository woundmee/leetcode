package main

func main() {
	findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
}

// problem --> https://neetcode.io/problems/max-consecutive-ones/

// Input: nums = [1,1,0,1,1,1]
// Output: 3

// вернуть максимальное кол-во последовательных 1

func findMaxConsecutiveOnes(nums []int) int {

	// [1,0,1,0,1,0,1]
	//    ^
	// res=1 count=0

	var max, count int

	for _, num := range nums {
		if num == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}
