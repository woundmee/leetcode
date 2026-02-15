package main

// problem --> https://neetcode.io/problems/implement-stack-using-queues/

func main() {

}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
	// slices.Reverse(this.queue)  // по условию задачи нельзя использовать

	// нужно перевернуть список
	// было: [1 2 3]
	// стало: [3 2 1]
	// т.е. очередь сделать ровном наоборот

	for range len(this.queue) - 1 {
		val := this.queue[0]                 // 1-ый элемент
		this.queue = this.queue[1:]          // удалили 1-ый элемент
		this.queue = append(this.queue, val) // добавили в конец

	}
}

func (this *MyStack) Pop() int {
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
