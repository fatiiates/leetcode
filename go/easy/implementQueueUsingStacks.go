package main

import "fmt"

type MyQueue struct {
	s1 MyStack
	s2 MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: ConstructorMyStack(),
		s2: ConstructorMyStack(),
	}
}

func (this *MyQueue) Push(x int) {
	for y := this.s1.Pop(); y != -1; y = this.s1.Pop() {
		this.s2.Push(y)
	}
	this.s1.Push(x)
	for y := this.s2.Pop(); y != -1; y = this.s2.Pop() {
		this.s1.Push(y)
	}
}

func (this *MyQueue) Pop() int {
	return this.s1.Pop()
}

func (this *MyQueue) Peek() int {
	return this.s1.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.Empty()
}

type MyStack struct {
	s []int
}

func ConstructorMyStack() MyStack {
	return MyStack{
		s: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.s = append(this.s, x)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	x := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	return x
}

func (this *MyStack) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.s[len(this.s)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.s) == 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Peek()
	param_2 := obj.Pop()
	param_4 := obj.Empty()
	fmt.Println(param_2, param_3, param_4)
	// obj := ConstructorMyStack()
	// obj.Push(1)
	// obj.Push(2)
	// fmt.Println(obj.Pop(), obj.Peek())
}
