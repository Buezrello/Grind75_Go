package solutions

type MyQueue struct {
	stack     []int
	tempStack []int
}

func Constructor() MyQueue {
	q := MyQueue{
		stack:     []int{},
		tempStack: []int{},
	}
	return q
}

func (this *MyQueue) Push(x int) {
	for !this.Empty() {
		popped := this.Pop()
		this.tempStack = append(this.tempStack, popped)
	}
	this.stack = append(this.stack, x)
	for len(this.tempStack) != 0 {
		popped := this.tempStack[len(this.tempStack)-1]
		this.tempStack = this.tempStack[:len(this.tempStack)-1]
		this.stack = append(this.stack, popped)
	}
}

func (this *MyQueue) Pop() int {
	popped := this.Peek()
	this.stack = this.stack[:len(this.stack)-1]
	return popped
}

func (this *MyQueue) Peek() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
