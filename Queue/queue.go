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

	)


type queue struct {
	ord []int
}

//representation of the queue, just the order
func (q queue) String() string {
	return fmt.Sprintf("%v", q.ord)
}

// add an integer to the back of the list
func (q *queue) Add(item int) []int {
	q.ord = append(q.ord, item)
	return q.ord
}

// remove
func (q *queue) Remove() int {
	output := q.ord[0]
	q.ord = q.ord[1:]
	return output
}

// look at the first position in the queue
func (q *queue) Peek() int {
	return q.ord[0]
}


func main(){
	q := queue{}
	q.Add(7)
	q.Add(8)
	q.Add(9)
	fmt.Println(q)
	fmt.Println(q.Peek())
	fmt.Println(q)
	fmt.Println(q.Remove())
	fmt.Println(q.Peek())
	fmt.Println(q)
}