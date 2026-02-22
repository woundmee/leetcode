package main

// problem --> https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/

func main() {
	// finalPrices([]int{8, 4, 6, 2, 3})
	finalPrices([]int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1})
}

// 8, 7, 4, 2, 8, 1, 7, 7, 10, 1
// 1  3  2  1  7

// ===================

// in: [8,4,6,2,3]
// out: [4,2,4,2,3]

// проходимся и находимся наименьший объект справа, далее prices[i]-minElement
// 8-4=4, 4-2=2, 6-2=4, 2, 3 - для у нет наименьшего правого - скидки нет.

func finalPrices(prices []int) []int {

	res := make([]int, len(prices))
	copy(res, prices)
	stack := []int{}

	for i, currentPrice := range prices {
		for len(stack) > 0 {
			lastIdx := stack[len(stack)-1]
			lastPrice := prices[lastIdx]
			if lastPrice >= currentPrice {
				res[lastIdx] = lastPrice - currentPrice
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return res
}
