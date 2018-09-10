/*
this contains the unit tests for the stack package

*/

package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	compare_stack := Stack{ord: []int{9, 1, 6}, min: 1}
	compare_stack2 := Stack{ord: []int{9, 1}, min: 1}
	s.Push(9)
	s.Push(1)
	s.Push(6)

	// test the Push method and structure
	if reflect.DeepEqual(s.ord, compare_stack.ord) != true {
		t.Errorf("Stacks not equal: %v, want: %v.", s, compare_stack)
	}
	// test the Min method
	if s.Min() != 1 {
		t.Errorf("Stacks min not equal: %v, want: 1.", s.Min())
	}
	// test the Pop method
	popped_val := s.Pop()
	if popped_val != 6 {
		t.Errorf("Stacks min not equal: %v, want: 6.", popped_val)
	}
	if reflect.DeepEqual(s.ord, compare_stack2.ord) != true {
		t.Errorf("Pop not working! Stacks not equal: %v, want: %v.", s, compare_stack2)
	}
	// test that Stringer is same as printing the ord
	print_string := fmt.Sprintf("%v", s)
	compare_print := fmt.Sprintf("%v", s.ord)
	if reflect.DeepEqual(print_string, compare_print) != true {
		t.Errorf("Stringer not working. Print strings not equal: %v, want: %v.", s, compare_stack2)
	}
}
