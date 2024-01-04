package main

import (
	"math"
)

type MinStack struct {
	array []int
	min   int
}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.array = append(this.array, this.min)
		this.min = x
	}
	this.array = append(this.array, x)
}

func (this *MinStack) Pop() {
	if len(this.array) == 0 {
		panic("empty stack")
	}

	if this.Top() == this.min {
		this.array = this.array[:len(this.array)-1]
		this.min = this.array[len(this.array)-1]
		this.array = this.array[:len(this.array)-1]
	} else {
		this.array = this.array[:len(this.array)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.array) == 0 {
		panic("empty stack")
	}
	return this.array[len(this.array)-1]
}

func (this *MinStack) Min() int {
	if len(this.array) == 0 {
		panic("empty stack")
	}
	return this.min
}

func main() {
}
