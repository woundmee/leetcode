package main

// problem --> https://neetcode.io/problems/implement-queue-using-stacks/

func main() {

}

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	// [1 2 3 4 5 ...]
	this.in = append(this.in, x) // просто добавляем элемент
}

// удаляет и возвращает этот же элемент
func (this *MyQueue) Pop() int {

	// логика такая же, как в Peek(), но с удалением
	peek := this.Peek()
	this.out = this.out[:len(this.out)-1]
	return peek
}

// просматриваем значение верхнего уровня
func (this *MyQueue) Peek() int {

	// in: 1 2 3
	// out 3 2 1

	if len(this.out) == 0 {
		// переложить in -> out
		for len(this.in) > 0 {
			lastIndex := len(this.in) - 1
			this.out = append(this.out, this.in[lastIndex])
			this.in = this.in[:lastIndex]
		}
	}

	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0 // если оба списка пустые - return true
}
