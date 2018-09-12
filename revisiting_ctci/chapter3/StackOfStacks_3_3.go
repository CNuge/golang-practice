/*

have a SetOfStacks function that will make a struct holding stacks,
when a given stack reaches a certain size, close it and then build a new
one so that they all stay small.

when popping stacks, need to close stacks when they become empty


*/

package main

import (
	"./stack"
	"fmt"
)

// a stack of stacks, where each stack will have a max of 10 items
type SetOfStacks []stack.Stack

func (s SetOfStacks) String() string {
	outstring := ""
	for i, dat := range s {
		outstring = fmt.Sprintf("%v%v\t%v\n", outstring, i, dat)
	}
	return outstring
}

func (s *SetOfStacks) Push(d int) {
	if len(*s) == 0 || len((*s)[len(*s)-1].Ord) >= 10 {
		new_stack := stack.Stack{}
		new_stack.Push(d)
		*s = append((*s), new_stack)
	} else {
		(*s)[len(*s)-1].Push(d)
	}

}

func (s *SetOfStacks) Pop() int {
	out := (*s)[len(*s)-1].Pop()

	// if the last stack is empty then close it
	if len((*s)[len(*s)-1].Ord) == 0 {
		*s = (*s)[0 : len(*s)-1]
	}
	return out
}

func (s *SetOfStacks) Peek() int {
	return (*s)[len(*s)-1].Peek()
}

func main() {

	test := SetOfStacks{}

	for i := 0; i < 22; i++ {
		test.Push(i)
	}

	fmt.Println(test)

	for i := 0; i < 3; i++ {
		x := test.Pop()
		fmt.Println(x)
	}

	fmt.Println(test)

}
