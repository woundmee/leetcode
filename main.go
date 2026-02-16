package main

// problem --> https://neetcode.io/problems/minimum-stack/

// stack -> LIFO

func main() {

}

type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		mins:  make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
	} else if this.mins[len(this.mins)-1] > val {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, this.mins[len(this.mins)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) != 0 && len(this.mins) != 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
