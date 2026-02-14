package main

// problem --> https://neetcode.io/problems/majority-element/question

func main() {
	_ = majorityElement([]int{5, 5, 1, 1, 1, 5, 5})
}

// вернуть элемент большинства из массива, который встречается n/2 раза, где n - длинна массива.

func majorityElement(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int)

	for i := range nums {
		if m[nums[i]] == nums[i] {
			m[nums[i]]++
			continue
		}

		m[nums[i]]++
	}

	max := 0
	val := 0
	for k, v := range m {
		if v > max {
			max = v
			val = k
		}
	}

	// fmt.Println(val)
	return val
}
