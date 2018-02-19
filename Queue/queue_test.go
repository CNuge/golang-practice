/*
queue is first in first out ordering.

code one with these operators:

	1. add - add to the back of the queue
	2. remove - remove the first in queue
	3. peek - look at the first object, but don't remove
	4. isempty - is the queue empty


*/

package queue

import(
	"fmt"
	"testing"
	"reflect"
	)


func TestQueue(t *testing.T){
	q := Queue{}
	compare_q := Queue{ord : []int {7,8,9}}
	compare_q2 := Queue{ord : []int {8,9}}
	
	q.Add(7)
	q.Add(8)
	q.Add(9)

	if reflect.DeepEqual(q.ord, compare_q.ord) != true {
		t.Errorf("Adding to Queue incorrect: %v, want: %v.", q, compare_q)
	}
	print_output := fmt.Sprintf("%v", q)
	if print_output != "[7 8 9]"{
		t.Errorf("Output string incorrect: %v, want: [7 8 9]", print_output)
	}
	if q.Peek() != 7 {
		t.Errorf("Peek() incorrect: %v, want: 7", q.Peek())
	}
	if reflect.DeepEqual(q.ord, compare_q.ord) != true {
		t.Errorf("Peek changed the queue!\n %v, want: [7 8 9]", print_output)
	}
	x := q.Remove()
	if x != 7 {
		t.Errorf("Remove() incorrect: %v, want: 7", print_output)
	}
	if reflect.DeepEqual(q.ord, compare_q2.ord) != true {
		t.Errorf("Remove() didn't change the queue!\n %v, want: [8 9]", print_output)
	}
}

