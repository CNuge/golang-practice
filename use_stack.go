package main

import(
	"fmt"
	"./stack"
	)



func main() {
	test_stack := stack.Stack{}
	test_stack.Push(12)
	test_stack.Push(13)
	test_stack.Push(14)
	test_stack.Push(3)
	fmt.Println(test_stack)
	fmt.Println(test_stack.Min())
	test_stack.Push(17)
	fmt.Println(test_stack)
	fmt.Println(test_stack.Pop())
	fmt.Println(test_stack)
}