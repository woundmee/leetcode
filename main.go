package main

import "fmt"

// problem:	https://neetcode.io/problems/majority-element-ii/
// task:	вернуть новый массив с числа, которые повторяются в исходном
//			больше чем n/3, где n - длина массива.

// notes
// [5,2,3,2,2,2,2,5,5,5], n=10
// n/3 = 3.3
// out: [2, 5]

func main() {
	majorityElement([]int{5, 2, 3, 2, 2, 2, 2, 5, 5, 5})
	majorityElement([]int{1, 2, 3})
}

func majorityElement(nums []int) []int {
	var res []int
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++

		if m[v] == len(nums)/3+1 {
			res = append(res, v)
		}

	}

	fmt.Println(res)
	return res
}
