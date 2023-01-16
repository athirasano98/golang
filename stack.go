package main

import "fmt"

type Stack struct {
	item []int
}

func (stack *Stack) Push(data int) {
	stack.item = append(stack.item, data)

}
func (stack *Stack) Pop() int {
	if len(stack.item) == 0 {
		return 0
	}
	data := stack.item[len(stack.item)-1]
	stack.item = stack.item[0 : len(stack.item)-1]
	return data
}
func (stack *Stack) Top() int {
	if len(stack.item) == 0 {
		return 0
	}
	data := stack.item[len(stack.item)-1]
	return data
}
func (stack *Stack) isEmpty() bool {
	return len(stack.item) == 0
}

func main() {
	var stack Stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)
	fmt.Println(stack.item)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.item)
	fmt.Println(stack.Top())
	fmt.Println(stack.isEmpty())

}
