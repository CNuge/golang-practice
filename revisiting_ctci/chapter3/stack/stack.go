/*
implement a stack with the following methods:

	1. push
	2. pop
	3. peek

recall stack is last in first out

*/

package stack

import (
	"fmt"
)

type Stack struct {
	Ord []int
}

// representation of the Stack
func (s Stack) String() string {
	return fmt.Sprintf("%v", s.Ord)
}

// add integer to the top of the stack
func (s *Stack) Push(item int) []int {
	s.Ord = append(s.Ord, item)
	return s.Ord
}

func (s *Stack) IsEmpty() bool {
	if len(s.Ord) == 0{
		return true
	}
	return false
}

// pop the top integer in the stack
func (s *Stack) Pop() int {
	output := s.Ord[len(s.Ord)-1]
	s.Ord = s.Ord[0 : len(s.Ord)-1]
	return output
}


func (s *Stack) Peek() int {
	output := s.Ord[len(s.Ord)-1]
	return output
}




