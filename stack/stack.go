/*
implement a stack with the following methods:

	1. push
	2. pop
	3. min (should be O(1))

recall stack is last in first out

*/

package stack

import(
	"fmt"
)

type Stack struct {
	ord []int
	min int
}

// representation of the Stack
func (s Stack) String() string {
	return fmt.Sprintf("%v", s.ord)
}

// add integer to the top of the stack
func (s *Stack) Push(item int) []int {
	if len(s.ord) == 0{
		s.min = item
	}
	if s.min > item {
		s.min = item
	}
	s.ord = append(s.ord, item)
	return s.ord
}

// pop the top integer in the stack
func (s *Stack) Pop() int {
	output := s.ord[len(s.ord)-1]
	s.ord = s.ord[0:len(s.ord)-1]
	return output
}

// find the minimum integer in the stack
func (s *Stack) Min() int {
	return s.min
}

