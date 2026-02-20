package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// int                              out
// ============================================
// 1212#                            abl
// 8512#12#15#23#15#18#12#4         helloworld

// ПРАВИЛА
// =======
// Символы с «a» по «i» отображаются в числа от «1» до «9» соответственно
// Символы с «j» по «z» отображаются в числа от «10#» до «26#» соответственно

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	var res []byte

	for i := len(line) - 1; i >= 0; {
		var number int

		lastSymbol := line[i]
		if lastSymbol == '#' {

			tens := int(line[i-2] - '0')
			ones := int(line[i-1] - '0')
			number = tens*10 + ones

			i -= 3
		} else {
			number = int(line[i] - '0')
			i -= 1
		}

		letter := byte('a' + number - 1)
		res = append(res, letter)
	}

	slices.Reverse(res)
	fmt.Println(string(res))
}
