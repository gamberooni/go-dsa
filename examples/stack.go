package main

import (
	"fmt"

	"github.com/gamberooni/go-dsa/stack"
)

// stack can be used to reverse numbers
func main() {
	stack := stack.NewStack(5, []int{})
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
