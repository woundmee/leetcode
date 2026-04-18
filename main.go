package main

// problem:	https://neetcode.io/problems/daily-temperatures/
// level:	medium

// task:	дан массив температур. Вернуть такой же массив, где будут цифры, указывающие
//			сколько детей до потепления. Пример: [30,38,30,36,35,40,28], на выходе: [1 4 1 2 1 0 0]
//			Объяснение:	1) 1-ый день = 30, следующее потепление = 38, разница 1
//						2) 2-ой день = 38, следующее потепление = 40, разница 4 и тд.

func main() {
	dailyTemperatures([]int{30, 38, 30, 36, 35, 40, 28})
	dailyTemperatures([]int{22, 21, 20})
}

// [30,38,30,36,35,40,28]
//                    ^
//  [1 4 1 2 1 0 0]

func dailyTemperatures(temperatures []int) []int {
	steps := []int{}
	for i := range len(temperatures) {
		next := nextWarming(temperatures[i], temperatures[i+1:])
		steps = append(steps, next)
	}
	return steps
}

func nextWarming(currTemp int, temperatures []int) int {
	var count int
	for _, v := range temperatures {
		count++
		if currTemp < v {
			return count
		}
	}
	return 0
}
