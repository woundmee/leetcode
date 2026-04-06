package main

// problem:	https://neetcode.io/problems/max-water-container/
// level:	medium

// task:	Площадь = длина*ширина (S=L*W) -> для прямоугольных
//			для круга (бассейн) -> pi*r*r (r - радиус круга)
//			найти 2 линии, которые дают max(min(height[i], height[j]) * (j - i))
//
//
//

func main() {
	maxArea([]int{1, 7, 2, 5, 4, 7, 3, 6})
}

func maxArea(heights []int) int {
	var res int

	i, j := 0, len(heights)-1

	for i <= j {
		w := abs(j - i)
		h := min(heights[i], heights[j])
		s := h * w
		res = max(res, s)

		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
